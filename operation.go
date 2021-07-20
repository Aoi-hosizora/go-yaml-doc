package goapidoc

// =========
// Operation
// =========

// Operation represents an api operation information of Document.
type Operation struct {
	method  string
	route   string
	summary string

	desc          string
	operationId   string
	schemes       []string
	consumes      []string
	produces      []string
	tags          []string
	securities    []string
	secsScopes    map[string][]string
	deprecated    bool
	reqExample    interface{}
	externalDoc   *ExternalDoc
	additionalDoc string
	params        []*Param
	responses     []*Response
}

// NewOperation creates a default Operation with given arguments.
func NewOperation(method, route, summary string) *Operation {
	return &Operation{method: method, route: route, summary: summary}
}

// NewGetOperation creates a get Operation with given arguments.
func NewGetOperation(route, summary string) *Operation {
	return NewOperation(GET, route, summary)
}

// NewPutOperation creates a put Operation with given arguments.
func NewPutOperation(route, summary string) *Operation {
	return NewOperation(PUT, route, summary)
}

// NewPostOperation creates a post Operation with given arguments.
func NewPostOperation(route, summary string) *Operation {
	return NewOperation(POST, route, summary)
}

// NewDeleteOperation creates a delete Operation with given arguments.
func NewDeleteOperation(route, summary string) *Operation {
	return NewOperation(DELETE, route, summary)
}

// NewOptionsOperation creates an options Operation with given arguments.
func NewOptionsOperation(route, summary string) *Operation {
	return NewOperation(OPTIONS, route, summary)
}

// NewHeadOperation creates a head Operation with given arguments.
func NewHeadOperation(route, summary string) *Operation {
	return NewOperation(HEAD, route, summary)
}

// NewPatchOperation creates a patch Operation with given arguments.
func NewPatchOperation(route, summary string) *Operation {
	return NewOperation(PATCH, route, summary)
}

func (o *Operation) GetMethod() string                        { return o.method }
func (o *Operation) GetRoute() string                         { return o.route }
func (o *Operation) GetSummary() string                       { return o.summary }
func (o *Operation) GetDesc() string                          { return o.desc }
func (o *Operation) GetOperationId() string                   { return o.operationId }
func (o *Operation) GetSchemes() []string                     { return o.schemes }
func (o *Operation) GetConsumes() []string                    { return o.consumes }
func (o *Operation) GetProduces() []string                    { return o.produces }
func (o *Operation) GetTags() []string                        { return o.tags }
func (o *Operation) GetSecurities() []string                  { return o.securities }
func (o *Operation) GetSecuritiesScopes() map[string][]string { return o.secsScopes }
func (o *Operation) GetDeprecated() bool                      { return o.deprecated }
func (o *Operation) GetRequestExample() interface{}           { return o.reqExample }
func (o *Operation) GetExternalDoc() *ExternalDoc             { return o.externalDoc }
func (o *Operation) GetAdditionalDoc() string                 { return o.additionalDoc }
func (o *Operation) GetParams() []*Param                      { return o.params }
func (o *Operation) GetResponses() []*Response                { return o.responses }

// Method sets the method in Operation.
func (o *Operation) Method(method string) *Operation {
	o.method = method
	return o
}

// Route sets the route in Operation.
func (o *Operation) Route(route string) *Operation {
	o.route = route
	return o
}

// Summary sets the summary in Operation.
func (o *Operation) Summary(summary string) *Operation {
	o.summary = summary
	return o
}

// Desc sets the desc in Operation.
func (o *Operation) Desc(desc string) *Operation {
	o.desc = desc
	return o
}

// OperationId sets the operationId in Operation.
func (o *Operation) OperationId(operationId string) *Operation {
	o.operationId = operationId
	return o
}

// Schemes sets the whole schemes in Operation.
func (o *Operation) Schemes(schemes ...string) *Operation {
	o.schemes = schemes
	return o
}

// AddSchemes adds some schemes into Operation.
func (o *Operation) AddSchemes(schemes ...string) *Operation {
	o.schemes = append(o.schemes, schemes...)
	return o
}

// Consumes sets the whole consumes in Operation.
func (o *Operation) Consumes(consumes ...string) *Operation {
	o.consumes = consumes
	return o
}

// AddConsumes adds some consumes into Operation.
func (o *Operation) AddConsumes(consumes ...string) *Operation {
	o.consumes = append(o.consumes, consumes...)
	return o
}

// Produces sets the whole produces in Operation.
func (o *Operation) Produces(produces ...string) *Operation {
	o.produces = produces
	return o
}

// AddProduces adds some produces into Operation.
func (o *Operation) AddProduces(produces ...string) *Operation {
	o.produces = append(o.produces, produces...)
	return o
}

// Tags sets the whole tags in Operation.
func (o *Operation) Tags(tags ...string) *Operation {
	o.tags = tags
	return o
}

// AddTags adds some tags into Operation.
func (o *Operation) AddTags(tags ...string) *Operation {
	o.tags = append(o.tags, tags...)
	return o
}

// Securities sets the whole security requirements in Operation.
func (o *Operation) Securities(securities ...string) *Operation {
	o.securities = securities
	return o
}

// AddSecurities adds some security requirements into Operation.
func (o *Operation) AddSecurities(securities ...string) *Operation {
	o.securities = append(o.securities, securities...)
	return o
}

// SetSecurityScopes sets a security's scopes in Operation.
func (o *Operation) SetSecurityScopes(security string, scopes ...string) *Operation {
	if o.secsScopes == nil {
		o.secsScopes = make(map[string][]string)
	}
	o.secsScopes[security] = scopes
	return o
}

// Deprecated sets the deprecated in Operation.
func (o *Operation) Deprecated(deprecated bool) *Operation {
	o.deprecated = deprecated
	return o
}

// RequestExample sets the request example in Operation, this is only supported in API Blueprint.
func (o *Operation) RequestExample(reqExample interface{}) *Operation {
	o.reqExample = reqExample
	return o
}

// ExternalDoc sets the external documentation in Operation.
func (o *Operation) ExternalDoc(doc *ExternalDoc) *Operation {
	o.externalDoc = doc
	return o
}

// AdditionalDoc sets the additional document in Operation, this is only supported in API Blueprint.
func (o *Operation) AdditionalDoc(doc string) *Operation {
	o.additionalDoc = doc
	return o
}

// Params sets the whole params in Operation.
func (o *Operation) Params(params ...*Param) *Operation {
	o.params = params
	return o
}

// AddParams adds some params into Operation.
func (o *Operation) AddParams(params ...*Param) *Operation {
	o.params = append(o.params, params...)
	return o
}

// Responses sets the whole responses in Operation.
func (o *Operation) Responses(responses ...*Response) *Operation {
	o.responses = responses
	return o
}

// AddResponses adds some responses into Operation.
func (o *Operation) AddResponses(responses ...*Response) *Operation {
	o.responses = append(o.responses, responses...)
	return o
}

// ========
// Response
// ========

// Response represents an operation response information of Operation.
type Response struct {
	code int
	typ  string

	desc          string
	examples      []*ResponseExample
	headers       []*ResponseHeader
	additionalDoc string
}

// NewResponse creates a default Response with given arguments.
func NewResponse(code int, typ string) *Response {
	return &Response{code: code, typ: typ}
}

func (r *Response) GetCode() int                    { return r.code }
func (r *Response) GetType() string                 { return r.typ }
func (r *Response) GetDesc() string                 { return r.desc }
func (r *Response) GetExamples() []*ResponseExample { return r.examples }
func (r *Response) GetHeaders() []*ResponseHeader   { return r.headers }
func (r *Response) GetAdditionalDoc() string        { return r.additionalDoc }

// Code sets the code in Response.
func (r *Response) Code(code int) *Response {
	r.code = code
	return r
}

// Type sets the type in Response.
func (r *Response) Type(typ string) *Response {
	r.typ = typ
	return r
}

// Desc sets the desc in Response.
func (r *Response) Desc(desc string) *Response {
	r.desc = desc
	return r
}

// Examples sets the whole response examples in Response.
// TODO BREAK CHANGES
func (r *Response) Examples(examples ...*ResponseExample) *Response {
	r.examples = examples
	return r
}

// AddExamples add some response examples into Response.
func (r *Response) AddExamples(examples ...*ResponseExample) *Response {
	r.examples = append(r.examples, examples...)
	return r
}

// Headers sets the whole response headers in Response.
func (r *Response) Headers(headers ...*ResponseHeader) *Response {
	r.headers = headers
	return r
}

// AddHeaders add some response headers in Response.
func (r *Response) AddHeaders(headers ...*ResponseHeader) *Response {
	r.headers = append(r.headers, headers...)
	return r
}

// AdditionalDoc sets the additional document in Response, this is only supported in API Blueprint.
func (r *Response) AdditionalDoc(doc string) *Response {
	r.additionalDoc = doc
	return r
}

// ===============
// ResponseExample
// ===============

// ResponseExample represents a response example information of Response.
type ResponseExample struct {
	mime    string
	example interface{}
}

// NewResponseExample creates a default ResponseExample with given arguments.
func NewResponseExample(mime string, example interface{}) *ResponseExample {
	return &ResponseExample{mime: mime, example: example}
}

func (r *ResponseExample) GetMime() string         { return r.mime }
func (r *ResponseExample) GetExample() interface{} { return r.example }

// Mime sets the mime in ResponseExample.
func (r *ResponseExample) Mime(mime string) *ResponseExample {
	r.mime = mime
	return r
}

// Example sets the example in ResponseExample.
func (r *ResponseExample) Example(example interface{}) *ResponseExample {
	r.example = example
	return r
}

// ==============
// ResponseHeader
// ==============

// ResponseHeader represents a response header information of Response.
// TODO BREAK CHANGES
type ResponseHeader struct {
	name    string
	typ     string // primitive type
	desc    string
	example interface{}
}

// NewHeader creates a default Header with given arguments.
func NewResponseHeader(name, typ, desc string) *ResponseHeader {
	return &ResponseHeader{name: name, typ: typ, desc: desc}
}

func (h *ResponseHeader) GetName() string         { return h.name }
func (h *ResponseHeader) GetType() string         { return h.typ }
func (h *ResponseHeader) GetDesc() string         { return h.desc }
func (h *ResponseHeader) GetExample() interface{} { return h.example }

// Name sets the name in ResponseHeader.
func (h *ResponseHeader) Name(name string) *ResponseHeader {
	h.name = name
	return h
}

// Type sets the type in ResponseHeader.
func (h *ResponseHeader) Type(typ string) *ResponseHeader {
	h.typ = typ
	return h
}

// Desc sets the desc in ResponseHeader.
func (h *ResponseHeader) Desc(desc string) *ResponseHeader {
	h.desc = desc
	return h
}

// Example sets the example in ResponseHeader.
func (h *ResponseHeader) Example(example interface{}) *ResponseHeader {
	h.example = example
	return h
}

// =====
// Param
// =====

// Param represents a request parameter information of Operation.
type Param struct {
	name     string
	in       string
	typ      string // string number integer boolean array file
	required bool
	desc     string

	allowEmpty       bool
	defaul           interface{}
	example          interface{}
	pattern          string
	enum             []interface{}
	minLength        *int
	maxLength        *int
	minItems         *int
	maxItems         *int
	uniqueItems      bool
	collectionFormat string
	minimum          *float64
	maximum          *float64
	exclusiveMin     bool
	exclusiveMax     bool
	multipleOf       float64
	itemOption       *ItemOption
	xmlRepr          *XMLRepr
}

// NewParam creates a default Param with given arguments.
func NewParam(name, in, typ string, required bool, desc string) *Param {
	return &Param{name: name, in: in, typ: typ, required: required, desc: desc}
}

// NewPathParam creates a path Param with given arguments.
func NewPathParam(name, typ string, required bool, desc string) *Param {
	return NewParam(name, PATH, typ, required, desc)
}

// NewQueryParam creates a query Param with given arguments.
func NewQueryParam(name, typ string, required bool, desc string) *Param {
	return NewParam(name, QUERY, typ, required, desc)
}

// NewFormParam creates a form Param with given arguments.
func NewFormParam(name, typ string, required bool, desc string) *Param {
	return NewParam(name, FORM, typ, required, desc)
}

// NewBodyParam creates a body Param with given arguments.
func NewBodyParam(name, typ string, required bool, desc string) *Param {
	return NewParam(name, BODY, typ, required, desc)
}

// NewHeaderParam creates a header Param with given arguments.
func NewHeaderParam(name, typ string, required bool, desc string) *Param {
	return NewParam(name, HEADER, typ, required, desc)
}

func (p *Param) GetName() string             { return p.name }
func (p *Param) GetInLoc() string            { return p.in }
func (p *Param) GetType() string             { return p.typ }
func (p *Param) GetRequired() bool           { return p.required }
func (p *Param) GetDesc() string             { return p.desc }
func (p *Param) GetAllowEmpty() bool         { return p.allowEmpty }
func (p *Param) GetDefault() interface{}     { return p.defaul }
func (p *Param) GetExample() interface{}     { return p.example }
func (p *Param) GetPattern() string          { return p.pattern }
func (p *Param) GetEnum() []interface{}      { return p.enum }
func (p *Param) GetMinLength() *int          { return p.minLength }
func (p *Param) GetMaxLength() *int          { return p.maxLength }
func (p *Param) GetMinItems() *int           { return p.minItems }
func (p *Param) GetMaxItems() *int           { return p.maxItems }
func (p *Param) GetUniqueItems() bool        { return p.uniqueItems }
func (p *Param) GetCollectionFormat() string { return p.collectionFormat }
func (p *Param) GetMinimum() *float64        { return p.minimum }
func (p *Param) GetMaximum() *float64        { return p.maximum }
func (p *Param) GetExclusiveMin() bool       { return p.exclusiveMin }
func (p *Param) GetExclusiveMax() bool       { return p.exclusiveMax }
func (p *Param) GetMultipleOf() float64      { return p.multipleOf }
func (p *Param) GetItemOption() *ItemOption  { return p.itemOption }
func (p *Param) GetXMLRepr() *XMLRepr        { return p.xmlRepr }

// Name sets the name in Param.
func (p *Param) Name(name string) *Param {
	p.name = name
	return p
}

// InLoc sets the in-location in Param.
func (p *Param) InLoc(in string) *Param {
	p.in = in
	return p
}

// Type sets the type in Param.
func (p *Param) Type(typ string) *Param {
	p.typ = typ
	return p
}

// Required sets the required in Param.
func (p *Param) Required(required bool) *Param {
	p.required = required
	return p
}

// Desc sets the desc in Param.
func (p *Param) Desc(desc string) *Param {
	p.desc = desc
	return p
}

// AllowEmpty sets the allowEmpty in Param.
func (p *Param) AllowEmpty(allow bool) *Param {
	p.allowEmpty = allow
	return p
}

// Default sets the default in Param.
func (p *Param) Default(defaul interface{}) *Param {
	p.defaul = defaul
	return p
}

// Example sets the example in Param.
func (p *Param) Example(example interface{}) *Param {
	p.example = example
	return p
}

// Pattern sets the pattern in Param.
func (p *Param) Pattern(pattern string) *Param {
	p.pattern = pattern
	return p
}

// Enum sets the whole enums in Param.
func (p *Param) Enum(enums ...interface{}) *Param {
	p.enum = enums
	return p
}

// MinLength sets the minLength in Param.
func (p *Param) MinLength(min int) *Param {
	p.minLength = &min
	return p
}

// MaxLength sets the maxLength in Param.
func (p *Param) MaxLength(max int) *Param {
	p.maxLength = &max
	return p
}

// LengthRange sets the minLength and maxLength in Param.
// TODO BREAK CHANGES
func (p *Param) LengthRange(min, max int) *Param {
	p.minLength = &min
	p.maxLength = &max
	return p
}

// MinItems sets the minItems in Param.
func (p *Param) MinItems(min int) *Param {
	p.minItems = &min
	return p
}

// MaxItems sets the maxItems in Param.
func (p *Param) MaxItems(max int) *Param {
	p.maxItems = &max
	return p
}

// ItemsRange sets the minItems and maxItems in Param.
func (p *Param) ItemsRange(min, max int) *Param {
	p.minItems = &min
	p.maxItems = &max
	return p
}

// UniqueItems sets the uniqueItems in Param.
func (p *Param) UniqueItems(unique bool) *Param {
	p.uniqueItems = unique
	return p
}

// CollectionFormat sets the collectionFormat in Param.
func (p *Param) CollectionFormat(collectionFormat string) *Param {
	p.collectionFormat = collectionFormat
	return p
}

// Minimum sets the minimum in Param.
// TODO BREAK CHANGES
func (p *Param) Minimum(min float64) *Param {
	p.minimum = &min
	return p
}

// Maximum sets the maximum in Param.
// TODO BREAK CHANGES
func (p *Param) Maximum(max float64) *Param {
	p.maximum = &max
	return p
}

// ValueRange sets the minimum and maximum in Param.
// TODO BREAK CHANGES
func (p *Param) ValueRange(min, max float64) *Param {
	p.minimum = &min
	p.maximum = &max
	return p
}

// ExclusiveMin sets the exclusiveMin in Param.
func (p *Param) ExclusiveMin(exclusiveMin bool) *Param {
	p.exclusiveMin = exclusiveMin
	return p
}

// ExclusiveMax sets the exclusiveMax in Param.
func (p *Param) ExclusiveMax(exclusiveMax bool) *Param {
	p.exclusiveMax = exclusiveMax
	return p
}

// MultipleOf sets the multipleOf in Param.
func (p *Param) MultipleOf(multipleOf float64) *Param {
	p.multipleOf = multipleOf
	return p
}

// ItemOption sets the item option in Param, this is only supported in Swagger.
func (p *Param) ItemOption(itemOption *ItemOption) *Param {
	p.itemOption = itemOption
	return p
}

// XMLRepr sets the xml repr in Param, this is only supported in Swagger.
func (p *Param) XMLRepr(repr *XMLRepr) *Param {
	p.xmlRepr = repr
	return p
}