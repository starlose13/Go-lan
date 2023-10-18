// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/aiplatform/v1/schema/predict/instance/video_action_recognition.proto

package instance

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Prediction input format for Video Action Recognition.
type VideoActionRecognitionPredictionInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Google Cloud Storage location of the video on which to perform the
	// prediction.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// The MIME type of the content of the video. Only the following are
	// supported: video/mp4 video/avi video/quicktime
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// The beginning, inclusive, of the video's time segment on which to perform
	// the prediction. Expressed as a number of seconds as measured from the
	// start of the video, with "s" appended at the end. Fractions are allowed,
	// up to a microsecond precision.
	TimeSegmentStart string `protobuf:"bytes,3,opt,name=time_segment_start,json=timeSegmentStart,proto3" json:"time_segment_start,omitempty"`
	// The end, exclusive, of the video's time segment on which to perform
	// the prediction. Expressed as a number of seconds as measured from the
	// start of the video, with "s" appended at the end. Fractions are allowed,
	// up to a microsecond precision, and "inf" or "Infinity" is allowed, which
	// means the end of the video.
	TimeSegmentEnd string `protobuf:"bytes,4,opt,name=time_segment_end,json=timeSegmentEnd,proto3" json:"time_segment_end,omitempty"`
}

func (x *VideoActionRecognitionPredictionInstance) Reset() {
	*x = VideoActionRecognitionPredictionInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoActionRecognitionPredictionInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoActionRecognitionPredictionInstance) ProtoMessage() {}

func (x *VideoActionRecognitionPredictionInstance) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoActionRecognitionPredictionInstance.ProtoReflect.Descriptor instead.
func (*VideoActionRecognitionPredictionInstance) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDescGZIP(), []int{0}
}

func (x *VideoActionRecognitionPredictionInstance) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *VideoActionRecognitionPredictionInstance) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *VideoActionRecognitionPredictionInstance) GetTimeSegmentStart() string {
	if x != nil {
		return x.TimeSegmentStart
	}
	return ""
}

func (x *VideoActionRecognitionPredictionInstance) GetTimeSegmentEnd() string {
	if x != nil {
		return x.TimeSegmentEnd
	}
	return ""
}

var File_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDesc = []byte{
	0x0a, 0x51, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x28, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x64, 0x42, 0xea, 0x02, 0x0a, 0x36, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x2d,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x5a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x3b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0xaa, 0x02, 0x32, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0xca, 0x02, 0x32, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x5c, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x5c, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0xea, 0x02, 0x38, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x3a, 0x3a, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x3a, 0x3a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDescData = file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_goTypes = []interface{}{
	(*VideoActionRecognitionPredictionInstance)(nil), // 0: google.cloud.aiplatform.v1.schema.predict.instance.VideoActionRecognitionPredictionInstance
}
var file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_init()
}
func file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_init() {
	if File_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoActionRecognitionPredictionInstance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto = out.File
	file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_schema_predict_instance_video_action_recognition_proto_depIdxs = nil
}
