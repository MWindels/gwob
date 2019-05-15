/*
Package main shows how to use the 'gwob' package to parse geometry data from OBJ files.

See also: https://github.com/udhos/gwob
*/
package main

import (
	"fmt"
	"log"

	"github.com/udhos/gwob"
)

func main() {

	// Set options
	options := &gwob.ObjParserOptions{
		LogStats: true,
		Logger:   func(msg string) { fmt.Println(msg) },
	}

	// Load OBJ
	fileObj := "red_cube.obj"
	o, errObj := gwob.NewObjFromFile(fileObj, options)
	if errObj != nil {
		log.Printf("obj: parse error input=%s: %v", fileObj, errObj)
		return
	}

	fileMtl := o.Mtllib

	// Load material lib
	lib, errMtl := gwob.ReadMaterialLibFromFile(fileMtl, options)
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
