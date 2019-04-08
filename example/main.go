/*
Package main shows how to use the 'gwob' package to parse geometry data from OBJ files.

See also: https://github.com/udhos/gwob
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/udhos/gwob"
)

func main() {

	// Set options
	options := &gwob.ObjParserOptions{
		LogStats: true,
		Logger:   func(msg string) { fmt.Println(msg) },
	}

	// Open OBJ file
	fileObj := "red_cube.obj"
	inputObj, errOpen := os.Open(fileObj)
	if errOpen != nil {
		log.Printf("open obj: %s: %v", fileObj, errOpen)
		return
	}

	defer inputObj.Close()

	// Load OBJ
	o, errObj := gwob.NewObjFromReader(fileObj, inputObj, options)
	if errObj != nil {
		log.Printf("obj: parse error input=%s: %v", fileObj, errObj)
		return
	}

	// Open MTL file
	fileMtl := o.Mtllib
	inputMtl, errOpenMtl := os.Open(fileMtl)
	if errOpenMtl != nil {
		log.Printf("open mtl: %s: %v", fileMtl, errOpenMtl)
		return
	}

	defer inputMtl.Close()

	// Load material lib
	lib, errMtl := gwob.ReadMaterialLibFromReader(inputMtl, options)
	if errMtl != nil {
		log.Printf("mtl: parse error input=%s: %v", fileMtl, errMtl)
		return
	}

	// Scan OBJ groups
	for _, g := range o.Groups {

		mtl, found := lib.Lib[g.Usemtl]
		if found {
			log.Printf("obj=%s lib=%s group=%s material=%s MapKd=%s Kd=%v", fileObj, fileMtl, g.Name, g.Usemtl, mtl.MapKd, mtl.Kd)
			continue
		}

		log.Printf("obj=%s lib=%s group=%s material=%s NOT FOUND", fileObj, fileMtl, g.Name, g.Usemtl)
	}
}
