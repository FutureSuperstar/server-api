package controller

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"net/http"
	"server-api/controller/request"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/klauspost/compress/zstd"
)

func Decompress(c *gin.Context) {
	req := &request.ZstdRequest{}
	err := c.ShouldBind(req)
	resp := request.CommonResponse{Code: 200}
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
	} else {
		zsResp := &request.ZstdResponse{}
		err = decodeRequest(c, req, zsResp)
		if err != nil {
			resp.Code = 500
			resp.Message = err.Error()
		} else {
			err = decodeResponse(c, req, zsResp)
			if err != nil {
				resp.Code = 500
				resp.Message = err.Error()
			}
		}
		resp.Data = zsResp
	}
	c.JSON(http.StatusOK, resp)
}

func decodeResponse(c *gin.Context, req *request.ZstdRequest, zsResp *request.ZstdResponse) error {
	if len(req.Response) <= 0 {
		return nil
	}
	decoded, err := base64.StdEncoding.DecodeString(req.Response)
	resp := request.CommonResponse{
		Code: 200,
		Data: nil,
	}
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		return err
	}
	data, err := DeCompressWithZstd(decoded)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		return err
	}
	zsResp.Response = string(data)
	return nil
}

func decodeRequest(c *gin.Context, req *request.ZstdRequest, zsResp *request.ZstdResponse) error {
	if len(req.Request) <= 0 {
		return nil
	}
	decoded, err := base64.StdEncoding.DecodeString(req.Request)
	resp := request.CommonResponse{
		Code: 200,
		Data: nil,
	}
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		return err
	}
	data, err := DeCompressWithZstd(decoded)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		return err
	}
	zsResp.Request = string(data)
	return nil
}

var decoderPool = sync.Pool{
	New: func() interface{} {
		dec, err := zstd.NewReader(nil)
		if err != nil {
			log.Fatalf("Failed to create new Zstd Decoder: %v", err)
		}
		return dec
	},
}

// DeCompressWithZstd zstd 解压,空字符串返回空字符串
func DeCompressWithZstd(compressedData []byte) ([]byte, error) {
	if len(compressedData) == 0 {
		return compressedData, errors.New("compressedData is empty")
	}

	dec := decoderPool.Get().(*zstd.Decoder)
	defer decoderPool.Put(dec)

	var decompressedData bytes.Buffer
	dec.Reset(bytes.NewReader(compressedData))

	_, err := io.Copy(&decompressedData, dec)
	if err != nil {
		return nil, err
	}

	return decompressedData.Bytes(), nil
}
