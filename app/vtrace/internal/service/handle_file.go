package service

import (
	"bytes"
	"io"

	v1 "gitee.com/qciip-icp/v-trace/api/vtrace/v1"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
)

// Upload 上传文件
func (s *VTraceService) Upload(stream v1.VTraceInterface_UploadServer) error {

	buf := bytes.Buffer{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			mds := make(map[string]string)
			for _, md := range req.GetMds() {
				mds[md.Key] = md.Value
			}

			url, err := s.cas.Store(stream.Context(), buf.Bytes(), mds)
			if err != nil {
				logger.Errorf("store file content failed: %v", err)

				return err
			}
			return stream.SendAndClose(&v1.UploadResponse{
				Url: url,
			})

		}
		if err != nil {
			logger.Errorf("receive file stream failed: %v", err)
			return err
		}
		buf.Write(req.GetContent())
	}

}
