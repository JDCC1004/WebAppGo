package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func getRandomImg(n int) []ImageData{
	var imagesData []ImageData
	images := getImgDir("Imagenes")

	rand.Seed(time.Now().UnixNano())

	selected := make(map[int]bool)

	for i := 0; i < n && i < len(images);{
		index := rand.Intn(len(images))
		if selected[index]{
			continue
		}
		selected[index] = true

		imgBase64 := codificarBase64(images[index])
		imagesData = append(imagesData, ImageData{
			Base64: imgBase64,
			Name: images[index],
		})
		i ++
	}

	return imagesData
}

func getImgDir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil{
		panic(err)
	}
	
	var imagenes []string
	for _, file := range files{
		if !file.IsDir() && (file.Name()[len(file.Name()) -4:] == ".jpg" || file.Name()[len(file.Name()) -4:] == ".jpeg" || file.Name()[len(file.Name()) -4:] == ".png"){
			imagenes = append(imagenes, dir + "/" + file.Name())
			fmt.Println(file)
		}
	}

	return imagenes
}

func codificarBase64(imgPath string) string{
	file, err := os.Open(imgPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil{
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(fileBytes)
}