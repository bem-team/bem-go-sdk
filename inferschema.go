// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bem

import (
	"bytes"
	"context"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/bem-team/bem-go-sdk/internal/apiform"
	"github.com/bem-team/bem-go-sdk/internal/apijson"
	"github.com/bem-team/bem-go-sdk/internal/requestconfig"
	"github.com/bem-team/bem-go-sdk/option"
	"github.com/bem-team/bem-go-sdk/packages/respjson"
)

// Infer JSON Schemas from uploaded documents using AI.
//
// Upload a file (PDF, image, spreadsheet, email, etc.) and receive a
// general-purpose JSON Schema that captures the document's structure. The inferred
// schema can be used directly as the `outputSchema` when creating Transform
// functions.
//
// The schema is designed to be broadly applicable to documents of the same type,
// not just the specific file uploaded.
//
// InferSchemaService contains methods and other services that help with
// interacting with the bem API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferSchemaService] method instead.
type InferSchemaService struct {
	options []option.RequestOption
}

// NewInferSchemaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferSchemaService(opts ...option.RequestOption) (r InferSchemaService) {
	r = InferSchemaService{}
	r.options = opts
	return
}

// **Analyze a file and infer a JSON Schema from its contents.**
//
// Accepts a file via multipart form upload and uses Gemini to analyze the
// document, returning a description of its contents, an inferred JSON Schema
// capturing all extractable fields, and document classification metadata.
//
// The returned schema is designed to be reusable across many similar documents of
// the same type, not just the specific file uploaded. It can be used directly as
// the `outputSchema` when creating a Transform function.
//
// The endpoint also detects whether the file contains multiple bundled documents
// and classifies the content nature (textual, visual, audio, video, or mixed).
//
// ## Supported file types
//
// PDF, PNG, JPEG, HEIC, HEIF, WebP, CSV, XLS, XLSX, DOCX, JSON, HTML, XML, EML,
// plain text, WAV, MP3, M4A, MP4.
//
// ## File size limit
//
// Maximum file size is **20 MB**.
//
// ## Example
//
// ```bash
//
//	curl -X POST https://api.bem.ai/v3/infer-schema \
//	  -H "x-api-key: YOUR_API_KEY" \
//	  -F "file=@invoice.pdf"
//
// ```
func (r *InferSchemaService) New(ctx context.Context, body InferSchemaNewParams, opts ...option.RequestOption) (res *InferSchemaNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v3/infer-schema"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Response from the infer-schema endpoint.
type InferSchemaNewResponse struct {
	// Analysis result returned by the infer-schema endpoint.
	Analysis InferSchemaNewResponseAnalysis `json:"analysis" api:"required"`
	// Original filename of the uploaded file.
	Filename string `json:"filename" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analysis    respjson.Field
		Filename    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferSchemaNewResponse) RawJSON() string { return r.JSON.raw }
func (r *InferSchemaNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Analysis result returned by the infer-schema endpoint.
type InferSchemaNewResponseAnalysis struct {
	// Classification of the primary content. One of: `textual`, `visual`, `audio`,
	// `video`, `mixed`.
	ContentNature string `json:"contentNature" api:"required"`
	// MIME content type of the uploaded file.
	ContentType string `json:"contentType" api:"required"`
	// 2-3 sentence description of what the file contains.
	Description string `json:"description" api:"required"`
	// List of distinct document types found in the file with counts.
	DocumentTypes []InferSchemaNewResponseAnalysisDocumentType `json:"documentTypes" api:"required"`
	// Original filename of the uploaded file.
	FileName string `json:"fileName" api:"required"`
	// High-level file category (e.g. "document", "image", "spreadsheet", "email").
	FileType string `json:"fileType" api:"required"`
	// Whether the file contains multiple separate documents bundled together.
	IsMultiDocument bool `json:"isMultiDocument" api:"required"`
	// Size of the uploaded file in bytes.
	SizeBytes int64 `json:"sizeBytes" api:"required"`
	// Inferred JSON Schema representing all extractable data fields.
	Schema any `json:"schema"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentNature   respjson.Field
		ContentType     respjson.Field
		Description     respjson.Field
		DocumentTypes   respjson.Field
		FileName        respjson.Field
		FileType        respjson.Field
		IsMultiDocument respjson.Field
		SizeBytes       respjson.Field
		Schema          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferSchemaNewResponseAnalysis) RawJSON() string { return r.JSON.raw }
func (r *InferSchemaNewResponseAnalysis) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a distinct document type found in the file.
type InferSchemaNewResponseAnalysisDocumentType struct {
	// Number of instances of this document type in the file.
	Count int64 `json:"count" api:"required"`
	// Brief description of this document type.
	Description string `json:"description" api:"required"`
	// Short snake_case name (e.g. "invoice", "receipt", "utility_bill").
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferSchemaNewResponseAnalysisDocumentType) RawJSON() string { return r.JSON.raw }
func (r *InferSchemaNewResponseAnalysisDocumentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferSchemaNewParams struct {
	// The file to analyze and infer a JSON schema from.
	File any `json:"file,omitzero" api:"required"`
	paramObj
}

func (r InferSchemaNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
