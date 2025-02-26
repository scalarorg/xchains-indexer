package events

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/rs/zerolog/log"
	indexerModels "github.com/scalarorg/data-models/indexer"
	"github.com/scalarorg/xchains-indexer/config"
	"github.com/scalarorg/xchains-indexer/db/models"
	"github.com/scalarorg/xchains-indexer/indexer"
	"gorm.io/gorm"
)

type BaseParser struct {
	Id      string
	Indexer *indexer.Indexer
}

func (p *BaseParser) Identifier() string {
	return p.Id
}

// Custom event parser
func (p *BaseParser) ParseBlockEvent(block *models.Block, event abci.Event, attributes []models.BlockEventAttribute, conf config.IndexConfig) (*any, error) {
	return nil, nil
}

func (p *BaseParser) IndexBlockEvent(parsedData *any, db *gorm.DB, block models.Block, blockEvent models.BlockEvent, blockEventAttrs []models.BlockEventAttribute, cfg config.IndexConfig) error {
	if parsedData != nil {
		config.Log.Debugf("[BaseParser] IndexBlockEvent# Type %T, parsedData: %++v", *parsedData, *parsedData)
		res := db.Create(*parsedData)
		return res.Error
	}
	return nil
}

func GetAllFields(obj interface{}) map[string]reflect.StructField {
	fields := make(map[string]reflect.StructField)
	// Get the type of the struct
	v := reflect.TypeOf(obj)
	// Ensure obj is a struct or pointer to a struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// Ensure it's a struct
	if v.Kind() == reflect.Struct {
		// Loop through struct fields
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			tag := field.Tag.Get("json")
			if tag == "" {
				tag = field.Name
			}
			fields[tag] = field
		}
	} else {
		fmt.Println("Provided value is not a struct")
	}

	return fields
}

// FindFieldsByType loops through struct fields and returns field names of a specific type
func FindAssetFields(obj interface{}) map[string]reflect.StructField {
	fields := make(map[string]reflect.StructField)

	// Get the type of the struct
	v := reflect.TypeOf(obj)
	// Ensure obj is a struct or pointer to a struct
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// Ensure it's a struct
	if v.Kind() == reflect.Struct {
		targetType := reflect.TypeOf(indexerModels.Asset{})
		// Loop through struct fields
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			// Check if field type matches targetType
			if field.Type == targetType {
				tag := field.Tag.Get("json")
				if tag == "" {
					tag = field.Name
				}
				fields[tag] = field
			}
		}
	} else {
		fmt.Println("Provided value is not a struct")
	}

	return fields
}

func processAttribute[T any](msg *T, attribute models.BlockEventAttribute, allFields map[string]reflect.StructField, data map[string]string) {
	field, ok := allFields[attribute.BlockEventAttributeKey.Key]
	if ok {
		switch field.Type {
		case reflect.TypeOf(indexerModels.Asset{}):
			processAssetField(msg, attribute, field, data)
		case reflect.TypeOf([]byte{}):
			processBytesField(msg, attribute, field, data)
		default:
			data[attribute.BlockEventAttributeKey.Key] = strings.Trim(attribute.Value, "\"")
		}
	} else {
		data[attribute.BlockEventAttributeKey.Key] = strings.Trim(attribute.Value, "\"")
	}
}

func processAssetField[T any](msg *T, attribute models.BlockEventAttribute, field reflect.StructField, data map[string]string) {
	asset := &indexerModels.Asset{}
	if err := json.Unmarshal([]byte(attribute.Value), asset); err != nil {
		log.Error().Err(err).Any("JsonData", attribute.Value).Msg("Cannot unmarshal asset")
		return
	}

	data[attribute.BlockEventAttributeKey.Key+"_denom"] = strings.Trim(asset.Denom, "\"")
	amount, err := strconv.ParseUint(asset.Amount, 10, 64)
	if err != nil {
		log.Error().Err(err).Any("Asset", asset).Msg("Cannot parse asset amount")
		return
	}
	SetField(msg, field.Name+"Amount", amount)
}

func processBytesField[T any](msg *T, attribute models.BlockEventAttribute, field reflect.StructField, data map[string]string) {
	var bytea []byte
	value := strings.Trim(attribute.Value, "\"")
	// Try to unmarshal as a byte array first using format: [1,2,3]
	err := json.Unmarshal([]byte(value), &bytea)
	if err == nil {
		SetField(msg, field.Name, bytea)
	} else {
		// set default base64 encoded value
		data[attribute.BlockEventAttributeKey.Key] = value
	}
}
func CreateEvent[T any](block *models.Block, attributes []models.BlockEventAttribute) *T {
	msg := new(T)
	data := map[string]string{}
	//assetFields := FindAssetFields(msg)
	objectFields := GetAllFields(msg)
	for _, attribute := range attributes {
		processAttribute(msg, attribute, objectFields, data)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Any("JsonData", data).Msg("Cannot marshal message")
	}

	if err := json.Unmarshal(jsonData, msg); err != nil {
		log.Error().Err(err).Any("JsonData", data).Msg("Cannot unmarshal message")
	}

	// Set block info
	setBlockInfo(msg, block)
	return msg
}

func setBlockInfo[T any](msg *T, block *models.Block) {
	fields := map[string]interface{}{
		"BlockHeight": uint64(block.Height),
		"BlockHash":   block.Hash,
		"Timestamp":   block.TimeStamp,
	}

	for field, value := range fields {
		if err := SetField(msg, field, value); err != nil {
			log.Error().Err(err).Str("field", field).Msg("Cannot set block field")
		}
	}
}

// SetField dynamically sets a struct field if it exists and is settable
func SetField(obj interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(obj).Elem() // Get the value of the pointer
	f := v.FieldByName(fieldName)    // Get the field by name

	// Check if the field exists
	if !f.IsValid() {
		return fmt.Errorf("no such field: %s", fieldName)
	}

	// Check if the field can be set
	if !f.CanSet() {
		return fmt.Errorf("cannot set field: %s", fieldName)
	}

	// Ensure the value type matches the field type
	newValue := reflect.ValueOf(value)
	if newValue.Type() != f.Type() {
		return fmt.Errorf("type mismatch: expected %s but got %s", f.Type(), newValue.Type())
	}

	// Set the field value
	f.Set(newValue)
	return nil
}

func NormalizeHash(hash string) string {
	return strings.ToLower(strings.TrimPrefix(hash, "0x"))
}

func DecodeIntArrayToBytes(input string) ([]byte, error) {
	// Parse the input string as a slice of integers
	var intArray []int
	err := json.Unmarshal([]byte(input), &intArray)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input: %v", err)
	}
	// Convert the slice of integers to a slice of bytes
	byteArray := make([]byte, len(intArray))
	for i, v := range intArray {
		byteArray[i] = byte(v)
	}
	return byteArray, nil
}

func DecodeIntArrayToHexString(input string) (string, error) {
	byteArray, err := DecodeIntArrayToBytes(input)
	if err != nil {
		return "", err
	}
	hexString := hex.EncodeToString(byteArray)
	return hexString, nil
}
