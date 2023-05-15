package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Реализовать паттерн «адаптер» на любом примере.

// Структура, описывающая формат данных
type Human struct {
	Name string
	Age  int
}

// Структура service может хранить данные
type Service struct {
	data []byte
}

// addXMLData записывает xml в Service
func (s *Service) writeXMLData(xg XML) error { // Принимает объект, который реализует интерфейс XML
	data, err := xg.GetXML() 
	if err != nil {
		return err
	}
	s.data = data
	return nil
}

// Интерфейс получения данных XML
type XML interface {
	GetXML() ([]byte, error)
}

type xmlData struct {
	xmlDoc []byte
}

func (x *xmlData) GetXML() ([]byte, error) {
	return x.xmlDoc, nil
}

// Структура с данными Json
type jsonData struct {
	jsonDoc []byte
}

func (j *jsonData) GetJson() []byte {
	return j.jsonDoc
}

// Адаптер, реализующий XML
type jsonAdapter struct {
	*jsonData
}

// Вызывает GetJson у jsonData и возвращает XML
func (ja *jsonAdapter) GetXML() ([]byte, error) {
	var h Human
	jsonDoc := ja.GetJson()
	if err := json.Unmarshal(jsonDoc, &h); err != nil { // преобразовываем json в нашу структуру
		return nil, err
	}
	res, err := xml.Marshal(&h) // преобразовываем нашу структуру в XML
	if err != nil {
		return nil, err
	}
	return res, nil
}

func main() {

	service := &Service{}

	jd := &jsonData{
		[]byte(`{"Name":"Ira", "Age":10}`),
	}
	ja := &jsonAdapter{jd}

	service.writeXMLData(ja)
	fmt.Println(string(service.data))
}
