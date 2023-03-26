package concurrency

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type JsonObject struct {
	author     string
	createdUtc int
}

type JsonFile struct {
	fileName    string
	jsonObjects []JsonObject
}

func complex_example() {

	before := time.Now()

	// have workers read in all the files in the dataFiles directory
	err := os.Chdir("./go_warmup/Concurrency")
	if err != nil {
		return
	}
	fileInfos, _ := ioutil.ReadDir("./dataFiles")

	fileNameChannel := make(chan string, len(fileInfos))

	jsonFileChannel := make(chan JsonFile, len(fileInfos))

	for element := range fileInfos {
		fileNameChannel <- fileInfos[element].Name()
	}

	close(fileNameChannel)

	var firstGroup sync.WaitGroup
	var secondGroup sync.WaitGroup

	go func() {
		for jsonFile := range jsonFileChannel {
			jsonFile := jsonFile
			go func() {
				secondGroup.Add(1)
				makeJsonFile(jsonFile)
				secondGroup.Done()
			}()
		}
	}()

	for item := range fileNameChannel {
		item := item
		go func() {
			firstGroup.Add(1)
			processFile(item, jsonFileChannel)
			firstGroup.Done()
		}()
	}

	firstGroup.Wait()

	close(jsonFileChannel)

	secondGroup.Wait()

	after := time.Now()

	print("\nTime elapsed: " + (after.Sub(before).String()) + "\n")
}

func processFile(filename string, jsonFiles chan JsonFile) {
	print("Processing file: " + filename + " \n")
	file, err := os.Open("./dataFiles/" + filename)
	checkError(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var jsonFile = JsonFile{filename, make([]JsonObject, 0)}
	fileContentsChan := make(chan string)
	jsonObjectsChan := make(chan JsonObject)
	var group sync.WaitGroup
	group.Add(1)

	go func() {
		defer close(fileContentsChan)
		for scanner.Scan() {
			jsonString := scanner.Text()
			fileContentsChan <- jsonString
		}
		group.Done()
	}()

	go func() {
		group.Add(1)
		var jsonObjectsGroup sync.WaitGroup
		jsonObjectsGroup.Add(1)
		go func() {
			jsonObjectsGroup.Add(1)
			for {
				res, ok := <-fileContentsChan
				if ok == false {
					break
				} else {
					go func() {
						jsonObjectsGroup.Add(1)
						processJsonString(res, jsonObjectsChan)
						jsonObjectsGroup.Done()
					}()
				}
			}
			jsonObjectsGroup.Done()
		}()
		time.Sleep(20)
		jsonObjectsGroup.Done()
		jsonObjectsGroup.Wait()
		close(jsonObjectsChan)
		group.Done()
	}()

	var jsonGroup sync.WaitGroup

	jsonGroup.Add(1)
	go func() {
		for item := range jsonObjectsChan {
			item := item
			jsonFile.jsonObjects = append(jsonFile.jsonObjects, item)
		}
		jsonGroup.Done()
	}()

	group.Wait()

	err = file.Close()
	checkError(err)
	scanner = nil

	jsonGroup.Wait()

	jsonFiles <- jsonFile
}

func makeJsonFile(jsonFile JsonFile) {
	jsonObjects := jsonFile.jsonObjects
	file, err := os.Create("./processedFiles/" + jsonFile.fileName + ".json")
	if err != nil {
		return
	}
	w := bufio.NewWriter(file)
	for integer := range jsonObjects {
		jsonObject := jsonObjects[integer]
		_, _ = w.WriteString("{\"author\":\"" + jsonObject.author + "\",\"created_utc\":\"" + strconv.Itoa(jsonObject.createdUtc) + "\"}\n")
	}

	err = w.Flush()
	if err != nil {
		return
	}
	print("Wrote file: " + jsonFile.fileName + "\n")
	checkError(err)
	jsonFile.jsonObjects = nil
}

func processJsonString(jsonString string, jsonObjects chan JsonObject) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	checkError(err)
	var author string
	var createdDateTime int
	author = jsonMap["author"].(string)
	i := jsonMap["created_utc"]
	switch i.(type) {
	case string:
		createdDateTime, _ = strconv.Atoi(i.(string))
	case float64:
		createdDateTime = int(i.(float64))
	default:
		createdDateTime = 0
	}
	checkError(err)
	if author != "[deleted]" {
		js := JsonObject{author, createdDateTime}
		jsonObjects <- js
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
