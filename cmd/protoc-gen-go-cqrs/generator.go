package main

import (
	//"strings"

	"bufio"
	"fmt"
	"github.com/go-leo/cqrs/cmd/internal"
	"google.golang.org/protobuf/compiler/protogen"
	"os"
	"path/filepath"
	"strings"
)

func generateFile(gen *protogen.Plugin, file *protogen.File) {
	if len(file.Services) == 0 {
		return
	}
	generateFileContent(gen, file)
}

func generateFileContent(gen *protogen.Plugin, file *protogen.File) {
	if len(file.Services) == 0 {
		return
	}
	for _, service := range file.Services {
		files, err := getFileInfo(gen, file, service)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "warn: %s\n", err.Error())
			return
		}
		for _, f := range files {
			if err := f.Gen(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%s.%s error: %s \n", service.Desc.FullName(), f.Endpoint, err)
				continue
			}
		}
	}
}

func getFileInfo(gen *protogen.Plugin, file *protogen.File, service *protogen.Service) ([]*internal.File, error) {
	path := internal.NewPath(splitComment(service.Comments.Leading.String()))
	if len(path.Command) == 0 || len(path.Query) == 0 {
		return nil, fmt.Errorf(`%s QueryPath or CommandPath is empty`, service.Desc.FullName())
	}
	cwd, _ := os.Getwd()
	queryAbs := filepath.Join(filepath.Dir(filepath.Join(cwd, file.Desc.Path())), path.Query)
	commandAbs := filepath.Join(filepath.Dir(filepath.Join(cwd, file.Desc.Path())), path.Command)
	var files []*internal.File
	for _, method := range service.Methods {
		if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
			// Unary RPC method
			endpoint := method.GoName
			file := internal.NewFileFromComment(endpoint, queryAbs, commandAbs, splitComment(method.Comments.Leading.String()))
			if file == nil {
				continue
			}
			files = append(files, file)
		} else {
			// Streaming RPC method
			continue
		}
	}
	return files, nil
}

func splitComment(leadingComment string) []string {
	var comments []string
	scanner := bufio.NewScanner(strings.NewReader(leadingComment))
	for scanner.Scan() {
		line := scanner.Text()
		comments = append(comments, line)
	}
	return comments
}

func clientName(service *protogen.Service) string {
	return service.GoName + "Client"
}

func serverName(service *protogen.Service) string {
	return service.GoName + "Server"
}

func fullMethodName(service *protogen.Service, method *protogen.Method) string {
	return fmt.Sprintf("/%s/%s", service.Desc.FullName(), method.Desc.Name())
}