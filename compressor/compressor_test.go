package compressor

import (
	"fmt"
	"log"
	"testing"
)

func Test4Compressor(t *testing.T) {
	var compressData []byte
	var err error
	text := "The example we're going to use is a very simple \"address book\" application that can read and write people's contact details to and from a file. Each person in the address book has a name, an ID, an email address, and a contact phone number.\n\nHow do you serialize and retrieve structured data like this? There are a few ways to solve this problem:\n\nUse gobs to serialize Go data structures. This is a good solution in a Go-specific environment, but it doesn't work well if you need to share data with applications written for other platforms.\nYou can invent an ad-hoc way to encode the data items into a single string – such as encoding 4 ints as \"12:3:-23:67\". This is a simple and flexible approach, although it does require writing one-off encoding and parsing code, and the parsing imposes a small run-time cost. This works best for encoding very simple data.\nSerialize the data to XML. This approach can be very attractive since XML is (sort of) human readable and there are binding libraries for lots of languages. This can be a good choice if you want to share data with other applications/projects. However, XML is notoriously space intensive, and encoding/decoding it can impose a huge performance penalty on applications. Also, navigating an XML DOM tree is considerably more complicated than navigating simple fields in a class normally would be.\nProtocol buffers are the flexible, efficient, automated solution to solve exactly this problem. With protocol buffers, you write a .proto description of the data structure you wish to store. From that, the protocol buffer compiler creates a class that implements automatic encoding and parsing of the protocol buffer data with an efficient binary format. The generated class provides getters and setters for the fields that make up a protocol buffer and takes care of the details of reading and writing the protocol buffer as a unit. Importantly, the protocol buffer format supports the idea of extending the format over time in such a way that the code can still read data encoded with the old format."
	data := []byte(text)
	for id, compressor := range Compressors {
		if compressData, err = compressor.Zip(data); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("compressor id: %d, data size: %d, compressData size: %d\n", id, len(data), len(compressData))
	}
}
