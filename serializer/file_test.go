package serializer_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	pb_laptop "github.com/yangwawa0323/pcbook/pb/laptop/v1"
	"github.com/yangwawa0323/pcbook/sample"
	"github.com/yangwawa0323/pcbook/serializer"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb_laptop.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
	t.Logf("%v", laptop2)

	err = serializer.WriteProtobufToJSONFile(laptop2, jsonFile)
	require.NoError(t, err)

}
