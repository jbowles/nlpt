/*
Functions to traverse directories and get file names and stream input file through a tokenizer to an output file
*/
package tkz

import (
	"bufio"
	"bytes"
	"github.com/jbowles/nlpt_tkz/Godeps/_workspace/src/gopkg.in/pipe.v2"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// StreamTokenizedFile streams data from a specified file, tokenizes text on the stream and returns []byte output and error. Error should return nil and []bytes should be greater than one.
// Since we are only dealing with one file the byte size returned should not be huge and so we simply return the content for the user to handle.
func StreamTokenizedFile(wg *sync.WaitGroup, timeoutLimit time.Duration, tokenDelimter byte, inFile, outFile, tkzType string) {
	//overwrite the output file and close, so subsuequent runs don't append but pipes do
	f, ferr := os.Create(outFile)
	f.Close()
	if ferr != nil {
		wg.Done()
		panic(ferr)
	}
	//overwritten output file and closed it so we start fresh with Append

	p := pipe.Line(
		//PipeFileTokensOnePerLine(tokenizeAsBytes, inFile, tkzType),
		PipeFileTokens(tokenDelimter, inFile, tkzType),
		pipe.AppendFile(outFile, 0644),
	)

	_, err := pipe.CombinedOutputTimeout(p, timeoutLimit)
	if err != nil {
		panic(err)
	}

	log.Printf("pipe.Line stream done for input: '%s'   output: '%s'", inFile, outFile)
	wg.Done()
}

// StreamTokenizedDirectory will use a custom file handler to sequentially loop through a directory and stream each file trhough PipeFileTokens.
func StreamTokenizedDirectory(wg *sync.WaitGroup, timeoutLimit time.Duration, tokenDelimter byte, directoryPath, outFile, tkzType string) {
	//overwrite the output file and close, so subsuequent runs don't append but pipes do
	f, err := os.Create(outFile)
	f.Close()
	if err != nil {
		wg.Done()
		panic(err)
	}
	//overwritten output file and closed it so we start fresh with Append

	handler := NewFileHandler(directoryPath, tkzType)
	go func() {
		for idx, file := range handler.FullFilePaths {
			p := pipe.Line(
				//PipeFileTokensOnePerLine(tokenizeAsBytes, file, handler.Tokenizer),
				PipeFileTokens(tokenDelimter, file, handler.Tokenizer),
				pipe.AppendFile(outFile, 0644),
			)
			_, err := pipe.CombinedOutputTimeout(p, timeoutLimit)
			//output, err := pipe.CombinedOutputTimeout(p, timeoutLimit)
			if err != nil {
				panic(err)
			}

			/// *************** DEBUGGING ****************
			log.Printf("streaming %d/%d ...\n", idx, len(handler.FullFilePaths))
			//Log.Debug("FILE: %v\n filter %v\n", file, string(output))
			//Log.Debug("FILE: %v\n tokens %v\n", file, string(output))
			/// *************** DEBUGGING ****************
		}
		log.Printf("read %d files for directory %s", len(handler.FullFilePaths), handler.DirName)
		wg.Done()
	}()
}

// PipeFileTokens reads data from the file at path and writes it to the pipe's stdout one token per line so that when we write to a file its a "word" per line. I've hijacked the pipe projects ReadFile function and stuck a text tokenzer inside of it.
// The tokenizer used here MUST be 'lex' OR 'unicode'. The latter is the fastest but less flexible and comprehensive, while the former is not much slower it will return alot of symbols and punctuation. If all you need is "words" then use the 'unicode' tokenizer.
func PipeFileTokens(delim byte, readFile, tokenizer string) pipe.Pipe {
	//so we don't fail becuase of bad tokenizer input
	var tkzType string
	switch tokenizer {
	case "unicode":
		tkzType = tokenizer
	default:
		tkzType = "lex"
	}
	//Log.Debug("Using tokenizer type: %s", tkzType)

	return pipe.TaskFunc(func(s *pipe.State) error {
		file, err := os.Open(s.Path(readFile))
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(file)
		bufferCache := new(bytes.Buffer)
		byteLining := []byte{'\n'} //newline padding bytes for writing to file
		for scanner.Scan() {
			bufferCache.Write(
				TokenizeBytes(append(scanner.Bytes(), delim), tkzType).Bytes,
			)
			bufferCache.Write(byteLining)
			//follow each buffer write with a new line
		}

		//close file as soon as we can but no sooner.
		file.Close()
		//Log.Debug("streamBytes from tokenzier: %d", bufferCache.Len())
		_, err = io.Copy(s.Stdout, bufferCache)
		if err != nil {
			panic(err)
		}
		return err
	})
}

// FileHandler contains the directory path, list of file paths, and function to create full file paths.
type FileHandler struct {
	DirName       string
	DirPath       string
	DocumentLabel string
	Tokenizer     string
	FullFilePaths []string
	FileInfo      []os.FileInfo
	FullPathFn    func(string, string, string) string
}

var separator = string(filepath.Separator)

func NewDirHandler(dirPath, dirLabel, tokenizer string) *FileHandler {
	handler := &FileHandler{
		DirName:       dirPath + separator,
		DirPath:       dirPath,
		Tokenizer:     tokenizer,
		DocumentLabel: dirLabel,
		FullPathFn:    func(dirpath, sep, filename string) string { return dirpath + sep + filename },
	}
	handler.setFileNames()
	return handler
}

func NewFileHandler(dirPath, tokenizer string) *FileHandler {
	handler := &FileHandler{
		DirName:    dirPath + separator,
		DirPath:    dirPath,
		Tokenizer:  tokenizer,
		FullPathFn: func(dirpath, sep, filename string) string { return dirpath + sep + filename },
	}
	handler.setFileNames()
	return handler
}

func (handle *FileHandler) setFileNames() {
	handle.getFileInfo()
	//Log.Debug("number of files %d:", len(handle.FileInfo))
	for _, file := range handle.FileInfo {
		if file.Mode().IsRegular() {
			handle.FullFilePaths = append(
				handle.FullFilePaths,
				handle.FullPathFn(handle.DirPath, separator, file.Name()),
			)
		}
	}
}

func (handle *FileHandler) getFileInfo() {
	//Log.Debug("GetFileInfo for new FileHandler %s:", handle.DirPath)
	d, err := os.Open(handle.DirPath)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		panic(err)
	}
	handle.FileInfo = files
}

func (handle *FileHandler) FileByteSize() map[string]int64 {
	fbs := make(map[string]int64)
	for _, file := range handle.FileInfo {
		if file.Mode().IsRegular() {
			fbs[file.Name()] = file.Size()
		}
	}
	return fbs
}
