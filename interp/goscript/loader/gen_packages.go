
// This is generate file
package loader

import (
	"reflect"
	
	"bytes"
	"crypto/md5"
	"net/http"
	"regexp"
	"unicode/utf16"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"strings"
	"unicode"
	"sync/atomic"
	"encoding/json"
	"html/template"
	"path"
	"github.com/linkxzhou/gongx/packages/server"
	"mime/multipart"
	"crypto/tls"
	"crypto/sha1"
	"container/heap"
	"container/list"
	"encoding/base64"
	"encoding/hex"
	"html"
	"math/rand"
	"time"
	"net/url"
	"io"
	"container/ring"
	"sort"
	"fmt"
	"strconv"
	"unicode/utf8"
	"io/ioutil"
	"encoding/xml"
	"math"
	"sync"
	"encoding/binary"
)

func init() {
	
	RegisterPackage(&Package{
		Name:       "bytes",
		Path:       "bytes",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Buffer": reflect.TypeOf(func(bytes.Buffer){}).In(0),
			"Reader": reflect.TypeOf(func(bytes.Reader){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"ErrTooLarge": reflect.ValueOf(func (error){}),

		},
		Funcs: map[string]reflect.Value{
			"Compare": reflect.ValueOf(bytes.Compare),
			"Contains": reflect.ValueOf(bytes.Contains),
			"ContainsAny": reflect.ValueOf(bytes.ContainsAny),
			"ContainsRune": reflect.ValueOf(bytes.ContainsRune),
			"Count": reflect.ValueOf(bytes.Count),
			"Cut": reflect.ValueOf(bytes.Cut),
			"Equal": reflect.ValueOf(bytes.Equal),
			"EqualFold": reflect.ValueOf(bytes.EqualFold),
			"Fields": reflect.ValueOf(bytes.Fields),
			"FieldsFunc": reflect.ValueOf(bytes.FieldsFunc),
			"HasPrefix": reflect.ValueOf(bytes.HasPrefix),
			"HasSuffix": reflect.ValueOf(bytes.HasSuffix),
			"Index": reflect.ValueOf(bytes.Index),
			"IndexAny": reflect.ValueOf(bytes.IndexAny),
			"IndexByte": reflect.ValueOf(bytes.IndexByte),
			"IndexFunc": reflect.ValueOf(bytes.IndexFunc),
			"IndexRune": reflect.ValueOf(bytes.IndexRune),
			"Join": reflect.ValueOf(bytes.Join),
			"LastIndex": reflect.ValueOf(bytes.LastIndex),
			"LastIndexAny": reflect.ValueOf(bytes.LastIndexAny),
			"LastIndexByte": reflect.ValueOf(bytes.LastIndexByte),
			"LastIndexFunc": reflect.ValueOf(bytes.LastIndexFunc),
			"Map": reflect.ValueOf(bytes.Map),
			"NewBuffer": reflect.ValueOf(bytes.NewBuffer),
			"NewBufferString": reflect.ValueOf(bytes.NewBufferString),
			"NewReader": reflect.ValueOf(bytes.NewReader),
			"Repeat": reflect.ValueOf(bytes.Repeat),
			"Replace": reflect.ValueOf(bytes.Replace),
			"ReplaceAll": reflect.ValueOf(bytes.ReplaceAll),
			"Runes": reflect.ValueOf(bytes.Runes),
			"Split": reflect.ValueOf(bytes.Split),
			"SplitAfter": reflect.ValueOf(bytes.SplitAfter),
			"SplitAfterN": reflect.ValueOf(bytes.SplitAfterN),
			"SplitN": reflect.ValueOf(bytes.SplitN),
			"Title": reflect.ValueOf(bytes.Title),
			"ToLower": reflect.ValueOf(bytes.ToLower),
			"ToLowerSpecial": reflect.ValueOf(bytes.ToLowerSpecial),
			"ToTitle": reflect.ValueOf(bytes.ToTitle),
			"ToTitleSpecial": reflect.ValueOf(bytes.ToTitleSpecial),
			"ToUpper": reflect.ValueOf(bytes.ToUpper),
			"ToUpperSpecial": reflect.ValueOf(bytes.ToUpperSpecial),
			"ToValidUTF8": reflect.ValueOf(bytes.ToValidUTF8),
			"Trim": reflect.ValueOf(bytes.Trim),
			"TrimFunc": reflect.ValueOf(bytes.TrimFunc),
			"TrimLeft": reflect.ValueOf(bytes.TrimLeft),
			"TrimLeftFunc": reflect.ValueOf(bytes.TrimLeftFunc),
			"TrimPrefix": reflect.ValueOf(bytes.TrimPrefix),
			"TrimRight": reflect.ValueOf(bytes.TrimRight),
			"TrimRightFunc": reflect.ValueOf(bytes.TrimRightFunc),
			"TrimSpace": reflect.ValueOf(bytes.TrimSpace),
			"TrimSuffix": reflect.ValueOf(bytes.TrimSuffix),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "md5",
		Path:       "crypto/md5",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(md5.New),
			"Sum": reflect.ValueOf(md5.Sum),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "http",
		Path:       "net/http",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Client": reflect.TypeOf(func(http.Client){}).In(0),
			"CloseNotifier": reflect.TypeOf(func(http.CloseNotifier){}).In(0),
			"ConnState": reflect.TypeOf(func(http.ConnState){}).In(0),
			"Cookie": reflect.TypeOf(func(http.Cookie){}).In(0),
			"CookieJar": reflect.TypeOf(func(http.CookieJar){}).In(0),
			"Dir": reflect.TypeOf(func(http.Dir){}).In(0),
			"File": reflect.TypeOf(func(http.File){}).In(0),
			"FileSystem": reflect.TypeOf(func(http.FileSystem){}).In(0),
			"Flusher": reflect.TypeOf(func(http.Flusher){}).In(0),
			"Handler": reflect.TypeOf(func(http.Handler){}).In(0),
			"HandlerFunc": reflect.TypeOf(func(http.HandlerFunc){}).In(0),
			"Header": reflect.TypeOf(func(http.Header){}).In(0),
			"Hijacker": reflect.TypeOf(func(http.Hijacker){}).In(0),
			"ProtocolError": reflect.TypeOf(func(http.ProtocolError){}).In(0),
			"PushOptions": reflect.TypeOf(func(http.PushOptions){}).In(0),
			"Pusher": reflect.TypeOf(func(http.Pusher){}).In(0),
			"Request": reflect.TypeOf(func(http.Request){}).In(0),
			"Response": reflect.TypeOf(func(http.Response){}).In(0),
			"ResponseWriter": reflect.TypeOf(func(http.ResponseWriter){}).In(0),
			"RoundTripper": reflect.TypeOf(func(http.RoundTripper){}).In(0),
			"SameSite": reflect.TypeOf(func(http.SameSite){}).In(0),
			"ServeMux": reflect.TypeOf(func(http.ServeMux){}).In(0),
			"Server": reflect.TypeOf(func(http.Server){}).In(0),
			"Transport": reflect.TypeOf(func(http.Transport){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"DefaultClient": reflect.ValueOf(http.DefaultClient),
			"DefaultServeMux": reflect.ValueOf(http.DefaultServeMux),
			"ErrAbortHandler": reflect.ValueOf(func (error){}),
			"ErrBodyNotAllowed": reflect.ValueOf(func (error){}),
			"ErrBodyReadAfterClose": reflect.ValueOf(func (error){}),
			"ErrContentLength": reflect.ValueOf(func (error){}),
			"ErrHandlerTimeout": reflect.ValueOf(func (error){}),
			"ErrHeaderTooLong": reflect.ValueOf(http.ErrHeaderTooLong),
			"ErrHijacked": reflect.ValueOf(func (error){}),
			"ErrLineTooLong": reflect.ValueOf(func (error){}),
			"ErrMissingBoundary": reflect.ValueOf(http.ErrMissingBoundary),
			"ErrMissingContentLength": reflect.ValueOf(http.ErrMissingContentLength),
			"ErrMissingFile": reflect.ValueOf(func (error){}),
			"ErrNoCookie": reflect.ValueOf(func (error){}),
			"ErrNoLocation": reflect.ValueOf(func (error){}),
			"ErrNotMultipart": reflect.ValueOf(http.ErrNotMultipart),
			"ErrNotSupported": reflect.ValueOf(http.ErrNotSupported),
			"ErrServerClosed": reflect.ValueOf(func (error){}),
			"ErrShortBody": reflect.ValueOf(http.ErrShortBody),
			"ErrSkipAltProtocol": reflect.ValueOf(func (error){}),
			"ErrUnexpectedTrailer": reflect.ValueOf(http.ErrUnexpectedTrailer),
			"ErrUseLastResponse": reflect.ValueOf(func (error){}),
			"ErrWriteAfterFlush": reflect.ValueOf(func (error){}),
			"LocalAddrContextKey": reflect.ValueOf(http.LocalAddrContextKey),
			"NoBody": reflect.ValueOf(http.NoBody),
			"ServerContextKey": reflect.ValueOf(http.ServerContextKey),

		},
		Funcs: map[string]reflect.Value{
			"AllowQuerySemicolons": reflect.ValueOf(http.AllowQuerySemicolons),
			"CanonicalHeaderKey": reflect.ValueOf(http.CanonicalHeaderKey),
			"DetectContentType": reflect.ValueOf(http.DetectContentType),
			"Error": reflect.ValueOf(http.Error),
			"FS": reflect.ValueOf(http.FS),
			"FileServer": reflect.ValueOf(http.FileServer),
			"Get": reflect.ValueOf(http.Get),
			"Handle": reflect.ValueOf(http.Handle),
			"HandleFunc": reflect.ValueOf(http.HandleFunc),
			"Head": reflect.ValueOf(http.Head),
			"ListenAndServe": reflect.ValueOf(http.ListenAndServe),
			"ListenAndServeTLS": reflect.ValueOf(http.ListenAndServeTLS),
			"MaxBytesHandler": reflect.ValueOf(http.MaxBytesHandler),
			"MaxBytesReader": reflect.ValueOf(http.MaxBytesReader),
			"NewFileTransport": reflect.ValueOf(http.NewFileTransport),
			"NewRequest": reflect.ValueOf(http.NewRequest),
			"NewRequestWithContext": reflect.ValueOf(http.NewRequestWithContext),
			"NewServeMux": reflect.ValueOf(http.NewServeMux),
			"NotFound": reflect.ValueOf(http.NotFound),
			"NotFoundHandler": reflect.ValueOf(http.NotFoundHandler),
			"ParseHTTPVersion": reflect.ValueOf(http.ParseHTTPVersion),
			"ParseTime": reflect.ValueOf(http.ParseTime),
			"Post": reflect.ValueOf(http.Post),
			"PostForm": reflect.ValueOf(http.PostForm),
			"ProxyFromEnvironment": reflect.ValueOf(http.ProxyFromEnvironment),
			"ProxyURL": reflect.ValueOf(http.ProxyURL),
			"ReadRequest": reflect.ValueOf(http.ReadRequest),
			"ReadResponse": reflect.ValueOf(http.ReadResponse),
			"Redirect": reflect.ValueOf(http.Redirect),
			"RedirectHandler": reflect.ValueOf(http.RedirectHandler),
			"Serve": reflect.ValueOf(http.Serve),
			"ServeContent": reflect.ValueOf(http.ServeContent),
			"ServeFile": reflect.ValueOf(http.ServeFile),
			"ServeTLS": reflect.ValueOf(http.ServeTLS),
			"SetCookie": reflect.ValueOf(http.SetCookie),
			"StatusText": reflect.ValueOf(http.StatusText),
			"StripPrefix": reflect.ValueOf(http.StripPrefix),
			"TimeoutHandler": reflect.ValueOf(http.TimeoutHandler),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "regexp",
		Path:       "regexp",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Regexp": reflect.TypeOf(func(regexp.Regexp){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"Compile": reflect.ValueOf(regexp.Compile),
			"CompilePOSIX": reflect.ValueOf(regexp.CompilePOSIX),
			"Match": reflect.ValueOf(regexp.Match),
			"MatchReader": reflect.ValueOf(regexp.MatchReader),
			"MatchString": reflect.ValueOf(regexp.MatchString),
			"MustCompile": reflect.ValueOf(regexp.MustCompile),
			"MustCompilePOSIX": reflect.ValueOf(regexp.MustCompilePOSIX),
			"QuoteMeta": reflect.ValueOf(regexp.QuoteMeta),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "utf16",
		Path:       "unicode/utf16",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"Decode": reflect.ValueOf(utf16.Decode),
			"DecodeRune": reflect.ValueOf(utf16.DecodeRune),
			"Encode": reflect.ValueOf(utf16.Encode),
			"EncodeRune": reflect.ValueOf(utf16.EncodeRune),
			"IsSurrogate": reflect.ValueOf(utf16.IsSurrogate),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "cipher",
		Path:       "crypto/cipher",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"AEAD": reflect.TypeOf(func(cipher.AEAD){}).In(0),
			"Block": reflect.TypeOf(func(cipher.Block){}).In(0),
			"BlockMode": reflect.TypeOf(func(cipher.BlockMode){}).In(0),
			"Stream": reflect.TypeOf(func(cipher.Stream){}).In(0),
			"StreamReader": reflect.TypeOf(func(cipher.StreamReader){}).In(0),
			"StreamWriter": reflect.TypeOf(func(cipher.StreamWriter){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"NewCBCDecrypter": reflect.ValueOf(cipher.NewCBCDecrypter),
			"NewCBCEncrypter": reflect.ValueOf(cipher.NewCBCEncrypter),
			"NewCFBDecrypter": reflect.ValueOf(cipher.NewCFBDecrypter),
			"NewCFBEncrypter": reflect.ValueOf(cipher.NewCFBEncrypter),
			"NewCTR": reflect.ValueOf(cipher.NewCTR),
			"NewGCM": reflect.ValueOf(cipher.NewGCM),
			"NewGCMWithNonceSize": reflect.ValueOf(cipher.NewGCMWithNonceSize),
			"NewGCMWithTagSize": reflect.ValueOf(cipher.NewGCMWithTagSize),
			"NewOFB": reflect.ValueOf(cipher.NewOFB),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "des",
		Path:       "crypto/des",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"KeySizeError": reflect.TypeOf(func(des.KeySizeError){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"NewCipher": reflect.ValueOf(des.NewCipher),
			"NewTripleDESCipher": reflect.ValueOf(des.NewTripleDESCipher),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "errors",
		Path:       "errors",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"As": reflect.ValueOf(errors.As),
			"Is": reflect.ValueOf(errors.Is),
			"New": reflect.ValueOf(errors.New),
			"Unwrap": reflect.ValueOf(errors.Unwrap),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "strings",
		Path:       "strings",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Builder": reflect.TypeOf(func(strings.Builder){}).In(0),
			"Reader": reflect.TypeOf(func(strings.Reader){}).In(0),
			"Replacer": reflect.TypeOf(func(strings.Replacer){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"Clone": reflect.ValueOf(strings.Clone),
			"Compare": reflect.ValueOf(strings.Compare),
			"Contains": reflect.ValueOf(strings.Contains),
			"ContainsAny": reflect.ValueOf(strings.ContainsAny),
			"ContainsRune": reflect.ValueOf(strings.ContainsRune),
			"Count": reflect.ValueOf(strings.Count),
			"Cut": reflect.ValueOf(strings.Cut),
			"EqualFold": reflect.ValueOf(strings.EqualFold),
			"Fields": reflect.ValueOf(strings.Fields),
			"FieldsFunc": reflect.ValueOf(strings.FieldsFunc),
			"HasPrefix": reflect.ValueOf(strings.HasPrefix),
			"HasSuffix": reflect.ValueOf(strings.HasSuffix),
			"Index": reflect.ValueOf(strings.Index),
			"IndexAny": reflect.ValueOf(strings.IndexAny),
			"IndexByte": reflect.ValueOf(strings.IndexByte),
			"IndexFunc": reflect.ValueOf(strings.IndexFunc),
			"IndexRune": reflect.ValueOf(strings.IndexRune),
			"Join": reflect.ValueOf(strings.Join),
			"LastIndex": reflect.ValueOf(strings.LastIndex),
			"LastIndexAny": reflect.ValueOf(strings.LastIndexAny),
			"LastIndexByte": reflect.ValueOf(strings.LastIndexByte),
			"LastIndexFunc": reflect.ValueOf(strings.LastIndexFunc),
			"Map": reflect.ValueOf(strings.Map),
			"NewReader": reflect.ValueOf(strings.NewReader),
			"NewReplacer": reflect.ValueOf(strings.NewReplacer),
			"Repeat": reflect.ValueOf(strings.Repeat),
			"Replace": reflect.ValueOf(strings.Replace),
			"ReplaceAll": reflect.ValueOf(strings.ReplaceAll),
			"Split": reflect.ValueOf(strings.Split),
			"SplitAfter": reflect.ValueOf(strings.SplitAfter),
			"SplitAfterN": reflect.ValueOf(strings.SplitAfterN),
			"SplitN": reflect.ValueOf(strings.SplitN),
			"Title": reflect.ValueOf(strings.Title),
			"ToLower": reflect.ValueOf(strings.ToLower),
			"ToLowerSpecial": reflect.ValueOf(strings.ToLowerSpecial),
			"ToTitle": reflect.ValueOf(strings.ToTitle),
			"ToTitleSpecial": reflect.ValueOf(strings.ToTitleSpecial),
			"ToUpper": reflect.ValueOf(strings.ToUpper),
			"ToUpperSpecial": reflect.ValueOf(strings.ToUpperSpecial),
			"ToValidUTF8": reflect.ValueOf(strings.ToValidUTF8),
			"Trim": reflect.ValueOf(strings.Trim),
			"TrimFunc": reflect.ValueOf(strings.TrimFunc),
			"TrimLeft": reflect.ValueOf(strings.TrimLeft),
			"TrimLeftFunc": reflect.ValueOf(strings.TrimLeftFunc),
			"TrimPrefix": reflect.ValueOf(strings.TrimPrefix),
			"TrimRight": reflect.ValueOf(strings.TrimRight),
			"TrimRightFunc": reflect.ValueOf(strings.TrimRightFunc),
			"TrimSpace": reflect.ValueOf(strings.TrimSpace),
			"TrimSuffix": reflect.ValueOf(strings.TrimSuffix),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "unicode",
		Path:       "unicode",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CaseRange": reflect.TypeOf(func(unicode.CaseRange){}).In(0),
			"Range16": reflect.TypeOf(func(unicode.Range16){}).In(0),
			"Range32": reflect.TypeOf(func(unicode.Range32){}).In(0),
			"RangeTable": reflect.TypeOf(func(unicode.RangeTable){}).In(0),
			"SpecialCase": reflect.TypeOf(func(unicode.SpecialCase){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"ASCII_Hex_Digit": reflect.ValueOf(unicode.ASCII_Hex_Digit),
			"Adlam": reflect.ValueOf(unicode.Adlam),
			"Ahom": reflect.ValueOf(unicode.Ahom),
			"Anatolian_Hieroglyphs": reflect.ValueOf(unicode.Anatolian_Hieroglyphs),
			"Arabic": reflect.ValueOf(unicode.Arabic),
			"Armenian": reflect.ValueOf(unicode.Armenian),
			"Avestan": reflect.ValueOf(unicode.Avestan),
			"AzeriCase": reflect.ValueOf(unicode.AzeriCase),
			"Balinese": reflect.ValueOf(unicode.Balinese),
			"Bamum": reflect.ValueOf(unicode.Bamum),
			"Bassa_Vah": reflect.ValueOf(unicode.Bassa_Vah),
			"Batak": reflect.ValueOf(unicode.Batak),
			"Bengali": reflect.ValueOf(unicode.Bengali),
			"Bhaiksuki": reflect.ValueOf(unicode.Bhaiksuki),
			"Bidi_Control": reflect.ValueOf(unicode.Bidi_Control),
			"Bopomofo": reflect.ValueOf(unicode.Bopomofo),
			"Brahmi": reflect.ValueOf(unicode.Brahmi),
			"Braille": reflect.ValueOf(unicode.Braille),
			"Buginese": reflect.ValueOf(unicode.Buginese),
			"Buhid": reflect.ValueOf(unicode.Buhid),
			"C": reflect.ValueOf(unicode.C),
			"Canadian_Aboriginal": reflect.ValueOf(unicode.Canadian_Aboriginal),
			"Carian": reflect.ValueOf(unicode.Carian),
			"CaseRanges": reflect.ValueOf(unicode.CaseRanges),
			"Categories": reflect.ValueOf(unicode.Categories),
			"Caucasian_Albanian": reflect.ValueOf(unicode.Caucasian_Albanian),
			"Cc": reflect.ValueOf(unicode.Cc),
			"Cf": reflect.ValueOf(unicode.Cf),
			"Chakma": reflect.ValueOf(unicode.Chakma),
			"Cham": reflect.ValueOf(unicode.Cham),
			"Cherokee": reflect.ValueOf(unicode.Cherokee),
			"Chorasmian": reflect.ValueOf(unicode.Chorasmian),
			"Co": reflect.ValueOf(unicode.Co),
			"Common": reflect.ValueOf(unicode.Common),
			"Coptic": reflect.ValueOf(unicode.Coptic),
			"Cs": reflect.ValueOf(unicode.Cs),
			"Cuneiform": reflect.ValueOf(unicode.Cuneiform),
			"Cypriot": reflect.ValueOf(unicode.Cypriot),
			"Cyrillic": reflect.ValueOf(unicode.Cyrillic),
			"Dash": reflect.ValueOf(unicode.Dash),
			"Deprecated": reflect.ValueOf(unicode.Deprecated),
			"Deseret": reflect.ValueOf(unicode.Deseret),
			"Devanagari": reflect.ValueOf(unicode.Devanagari),
			"Diacritic": reflect.ValueOf(unicode.Diacritic),
			"Digit": reflect.ValueOf(unicode.Digit),
			"Dives_Akuru": reflect.ValueOf(unicode.Dives_Akuru),
			"Dogra": reflect.ValueOf(unicode.Dogra),
			"Duployan": reflect.ValueOf(unicode.Duployan),
			"Egyptian_Hieroglyphs": reflect.ValueOf(unicode.Egyptian_Hieroglyphs),
			"Elbasan": reflect.ValueOf(unicode.Elbasan),
			"Elymaic": reflect.ValueOf(unicode.Elymaic),
			"Ethiopic": reflect.ValueOf(unicode.Ethiopic),
			"Extender": reflect.ValueOf(unicode.Extender),
			"FoldCategory": reflect.ValueOf(unicode.FoldCategory),
			"FoldScript": reflect.ValueOf(unicode.FoldScript),
			"Georgian": reflect.ValueOf(unicode.Georgian),
			"Glagolitic": reflect.ValueOf(unicode.Glagolitic),
			"Gothic": reflect.ValueOf(unicode.Gothic),
			"Grantha": reflect.ValueOf(unicode.Grantha),
			"GraphicRanges": reflect.ValueOf(unicode.GraphicRanges),
			"Greek": reflect.ValueOf(unicode.Greek),
			"Gujarati": reflect.ValueOf(unicode.Gujarati),
			"Gunjala_Gondi": reflect.ValueOf(unicode.Gunjala_Gondi),
			"Gurmukhi": reflect.ValueOf(unicode.Gurmukhi),
			"Han": reflect.ValueOf(unicode.Han),
			"Hangul": reflect.ValueOf(unicode.Hangul),
			"Hanifi_Rohingya": reflect.ValueOf(unicode.Hanifi_Rohingya),
			"Hanunoo": reflect.ValueOf(unicode.Hanunoo),
			"Hatran": reflect.ValueOf(unicode.Hatran),
			"Hebrew": reflect.ValueOf(unicode.Hebrew),
			"Hex_Digit": reflect.ValueOf(unicode.Hex_Digit),
			"Hiragana": reflect.ValueOf(unicode.Hiragana),
			"Hyphen": reflect.ValueOf(unicode.Hyphen),
			"IDS_Binary_Operator": reflect.ValueOf(unicode.IDS_Binary_Operator),
			"IDS_Trinary_Operator": reflect.ValueOf(unicode.IDS_Trinary_Operator),
			"Ideographic": reflect.ValueOf(unicode.Ideographic),
			"Imperial_Aramaic": reflect.ValueOf(unicode.Imperial_Aramaic),
			"Inherited": reflect.ValueOf(unicode.Inherited),
			"Inscriptional_Pahlavi": reflect.ValueOf(unicode.Inscriptional_Pahlavi),
			"Inscriptional_Parthian": reflect.ValueOf(unicode.Inscriptional_Parthian),
			"Javanese": reflect.ValueOf(unicode.Javanese),
			"Join_Control": reflect.ValueOf(unicode.Join_Control),
			"Kaithi": reflect.ValueOf(unicode.Kaithi),
			"Kannada": reflect.ValueOf(unicode.Kannada),
			"Katakana": reflect.ValueOf(unicode.Katakana),
			"Kayah_Li": reflect.ValueOf(unicode.Kayah_Li),
			"Kharoshthi": reflect.ValueOf(unicode.Kharoshthi),
			"Khitan_Small_Script": reflect.ValueOf(unicode.Khitan_Small_Script),
			"Khmer": reflect.ValueOf(unicode.Khmer),
			"Khojki": reflect.ValueOf(unicode.Khojki),
			"Khudawadi": reflect.ValueOf(unicode.Khudawadi),
			"L": reflect.ValueOf(unicode.L),
			"Lao": reflect.ValueOf(unicode.Lao),
			"Latin": reflect.ValueOf(unicode.Latin),
			"Lepcha": reflect.ValueOf(unicode.Lepcha),
			"Letter": reflect.ValueOf(unicode.Letter),
			"Limbu": reflect.ValueOf(unicode.Limbu),
			"Linear_A": reflect.ValueOf(unicode.Linear_A),
			"Linear_B": reflect.ValueOf(unicode.Linear_B),
			"Lisu": reflect.ValueOf(unicode.Lisu),
			"Ll": reflect.ValueOf(unicode.Ll),
			"Lm": reflect.ValueOf(unicode.Lm),
			"Lo": reflect.ValueOf(unicode.Lo),
			"Logical_Order_Exception": reflect.ValueOf(unicode.Logical_Order_Exception),
			"Lower": reflect.ValueOf(unicode.Lower),
			"Lt": reflect.ValueOf(unicode.Lt),
			"Lu": reflect.ValueOf(unicode.Lu),
			"Lycian": reflect.ValueOf(unicode.Lycian),
			"Lydian": reflect.ValueOf(unicode.Lydian),
			"M": reflect.ValueOf(unicode.M),
			"Mahajani": reflect.ValueOf(unicode.Mahajani),
			"Makasar": reflect.ValueOf(unicode.Makasar),
			"Malayalam": reflect.ValueOf(unicode.Malayalam),
			"Mandaic": reflect.ValueOf(unicode.Mandaic),
			"Manichaean": reflect.ValueOf(unicode.Manichaean),
			"Marchen": reflect.ValueOf(unicode.Marchen),
			"Mark": reflect.ValueOf(unicode.Mark),
			"Masaram_Gondi": reflect.ValueOf(unicode.Masaram_Gondi),
			"Mc": reflect.ValueOf(unicode.Mc),
			"Me": reflect.ValueOf(unicode.Me),
			"Medefaidrin": reflect.ValueOf(unicode.Medefaidrin),
			"Meetei_Mayek": reflect.ValueOf(unicode.Meetei_Mayek),
			"Mende_Kikakui": reflect.ValueOf(unicode.Mende_Kikakui),
			"Meroitic_Cursive": reflect.ValueOf(unicode.Meroitic_Cursive),
			"Meroitic_Hieroglyphs": reflect.ValueOf(unicode.Meroitic_Hieroglyphs),
			"Miao": reflect.ValueOf(unicode.Miao),
			"Mn": reflect.ValueOf(unicode.Mn),
			"Modi": reflect.ValueOf(unicode.Modi),
			"Mongolian": reflect.ValueOf(unicode.Mongolian),
			"Mro": reflect.ValueOf(unicode.Mro),
			"Multani": reflect.ValueOf(unicode.Multani),
			"Myanmar": reflect.ValueOf(unicode.Myanmar),
			"N": reflect.ValueOf(unicode.N),
			"Nabataean": reflect.ValueOf(unicode.Nabataean),
			"Nandinagari": reflect.ValueOf(unicode.Nandinagari),
			"Nd": reflect.ValueOf(unicode.Nd),
			"New_Tai_Lue": reflect.ValueOf(unicode.New_Tai_Lue),
			"Newa": reflect.ValueOf(unicode.Newa),
			"Nko": reflect.ValueOf(unicode.Nko),
			"Nl": reflect.ValueOf(unicode.Nl),
			"No": reflect.ValueOf(unicode.No),
			"Noncharacter_Code_Point": reflect.ValueOf(unicode.Noncharacter_Code_Point),
			"Number": reflect.ValueOf(unicode.Number),
			"Nushu": reflect.ValueOf(unicode.Nushu),
			"Nyiakeng_Puachue_Hmong": reflect.ValueOf(unicode.Nyiakeng_Puachue_Hmong),
			"Ogham": reflect.ValueOf(unicode.Ogham),
			"Ol_Chiki": reflect.ValueOf(unicode.Ol_Chiki),
			"Old_Hungarian": reflect.ValueOf(unicode.Old_Hungarian),
			"Old_Italic": reflect.ValueOf(unicode.Old_Italic),
			"Old_North_Arabian": reflect.ValueOf(unicode.Old_North_Arabian),
			"Old_Permic": reflect.ValueOf(unicode.Old_Permic),
			"Old_Persian": reflect.ValueOf(unicode.Old_Persian),
			"Old_Sogdian": reflect.ValueOf(unicode.Old_Sogdian),
			"Old_South_Arabian": reflect.ValueOf(unicode.Old_South_Arabian),
			"Old_Turkic": reflect.ValueOf(unicode.Old_Turkic),
			"Oriya": reflect.ValueOf(unicode.Oriya),
			"Osage": reflect.ValueOf(unicode.Osage),
			"Osmanya": reflect.ValueOf(unicode.Osmanya),
			"Other": reflect.ValueOf(unicode.Other),
			"Other_Alphabetic": reflect.ValueOf(unicode.Other_Alphabetic),
			"Other_Default_Ignorable_Code_Point": reflect.ValueOf(unicode.Other_Default_Ignorable_Code_Point),
			"Other_Grapheme_Extend": reflect.ValueOf(unicode.Other_Grapheme_Extend),
			"Other_ID_Continue": reflect.ValueOf(unicode.Other_ID_Continue),
			"Other_ID_Start": reflect.ValueOf(unicode.Other_ID_Start),
			"Other_Lowercase": reflect.ValueOf(unicode.Other_Lowercase),
			"Other_Math": reflect.ValueOf(unicode.Other_Math),
			"Other_Uppercase": reflect.ValueOf(unicode.Other_Uppercase),
			"P": reflect.ValueOf(unicode.P),
			"Pahawh_Hmong": reflect.ValueOf(unicode.Pahawh_Hmong),
			"Palmyrene": reflect.ValueOf(unicode.Palmyrene),
			"Pattern_Syntax": reflect.ValueOf(unicode.Pattern_Syntax),
			"Pattern_White_Space": reflect.ValueOf(unicode.Pattern_White_Space),
			"Pau_Cin_Hau": reflect.ValueOf(unicode.Pau_Cin_Hau),
			"Pc": reflect.ValueOf(unicode.Pc),
			"Pd": reflect.ValueOf(unicode.Pd),
			"Pe": reflect.ValueOf(unicode.Pe),
			"Pf": reflect.ValueOf(unicode.Pf),
			"Phags_Pa": reflect.ValueOf(unicode.Phags_Pa),
			"Phoenician": reflect.ValueOf(unicode.Phoenician),
			"Pi": reflect.ValueOf(unicode.Pi),
			"Po": reflect.ValueOf(unicode.Po),
			"Prepended_Concatenation_Mark": reflect.ValueOf(unicode.Prepended_Concatenation_Mark),
			"PrintRanges": reflect.ValueOf(unicode.PrintRanges),
			"Properties": reflect.ValueOf(unicode.Properties),
			"Ps": reflect.ValueOf(unicode.Ps),
			"Psalter_Pahlavi": reflect.ValueOf(unicode.Psalter_Pahlavi),
			"Punct": reflect.ValueOf(unicode.Punct),
			"Quotation_Mark": reflect.ValueOf(unicode.Quotation_Mark),
			"Radical": reflect.ValueOf(unicode.Radical),
			"Regional_Indicator": reflect.ValueOf(unicode.Regional_Indicator),
			"Rejang": reflect.ValueOf(unicode.Rejang),
			"Runic": reflect.ValueOf(unicode.Runic),
			"S": reflect.ValueOf(unicode.S),
			"STerm": reflect.ValueOf(unicode.STerm),
			"Samaritan": reflect.ValueOf(unicode.Samaritan),
			"Saurashtra": reflect.ValueOf(unicode.Saurashtra),
			"Sc": reflect.ValueOf(unicode.Sc),
			"Scripts": reflect.ValueOf(unicode.Scripts),
			"Sentence_Terminal": reflect.ValueOf(unicode.Sentence_Terminal),
			"Sharada": reflect.ValueOf(unicode.Sharada),
			"Shavian": reflect.ValueOf(unicode.Shavian),
			"Siddham": reflect.ValueOf(unicode.Siddham),
			"SignWriting": reflect.ValueOf(unicode.SignWriting),
			"Sinhala": reflect.ValueOf(unicode.Sinhala),
			"Sk": reflect.ValueOf(unicode.Sk),
			"Sm": reflect.ValueOf(unicode.Sm),
			"So": reflect.ValueOf(unicode.So),
			"Soft_Dotted": reflect.ValueOf(unicode.Soft_Dotted),
			"Sogdian": reflect.ValueOf(unicode.Sogdian),
			"Sora_Sompeng": reflect.ValueOf(unicode.Sora_Sompeng),
			"Soyombo": reflect.ValueOf(unicode.Soyombo),
			"Space": reflect.ValueOf(unicode.Space),
			"Sundanese": reflect.ValueOf(unicode.Sundanese),
			"Syloti_Nagri": reflect.ValueOf(unicode.Syloti_Nagri),
			"Symbol": reflect.ValueOf(unicode.Symbol),
			"Syriac": reflect.ValueOf(unicode.Syriac),
			"Tagalog": reflect.ValueOf(unicode.Tagalog),
			"Tagbanwa": reflect.ValueOf(unicode.Tagbanwa),
			"Tai_Le": reflect.ValueOf(unicode.Tai_Le),
			"Tai_Tham": reflect.ValueOf(unicode.Tai_Tham),
			"Tai_Viet": reflect.ValueOf(unicode.Tai_Viet),
			"Takri": reflect.ValueOf(unicode.Takri),
			"Tamil": reflect.ValueOf(unicode.Tamil),
			"Tangut": reflect.ValueOf(unicode.Tangut),
			"Telugu": reflect.ValueOf(unicode.Telugu),
			"Terminal_Punctuation": reflect.ValueOf(unicode.Terminal_Punctuation),
			"Thaana": reflect.ValueOf(unicode.Thaana),
			"Thai": reflect.ValueOf(unicode.Thai),
			"Tibetan": reflect.ValueOf(unicode.Tibetan),
			"Tifinagh": reflect.ValueOf(unicode.Tifinagh),
			"Tirhuta": reflect.ValueOf(unicode.Tirhuta),
			"Title": reflect.ValueOf(unicode.Title),
			"TurkishCase": reflect.ValueOf(unicode.TurkishCase),
			"Ugaritic": reflect.ValueOf(unicode.Ugaritic),
			"Unified_Ideograph": reflect.ValueOf(unicode.Unified_Ideograph),
			"Upper": reflect.ValueOf(unicode.Upper),
			"Vai": reflect.ValueOf(unicode.Vai),
			"Variation_Selector": reflect.ValueOf(unicode.Variation_Selector),
			"Wancho": reflect.ValueOf(unicode.Wancho),
			"Warang_Citi": reflect.ValueOf(unicode.Warang_Citi),
			"White_Space": reflect.ValueOf(unicode.White_Space),
			"Yezidi": reflect.ValueOf(unicode.Yezidi),
			"Yi": reflect.ValueOf(unicode.Yi),
			"Z": reflect.ValueOf(unicode.Z),
			"Zanabazar_Square": reflect.ValueOf(unicode.Zanabazar_Square),
			"Zl": reflect.ValueOf(unicode.Zl),
			"Zp": reflect.ValueOf(unicode.Zp),
			"Zs": reflect.ValueOf(unicode.Zs),

		},
		Funcs: map[string]reflect.Value{
			"In": reflect.ValueOf(unicode.In),
			"Is": reflect.ValueOf(unicode.Is),
			"IsControl": reflect.ValueOf(unicode.IsControl),
			"IsDigit": reflect.ValueOf(unicode.IsDigit),
			"IsGraphic": reflect.ValueOf(unicode.IsGraphic),
			"IsLetter": reflect.ValueOf(unicode.IsLetter),
			"IsLower": reflect.ValueOf(unicode.IsLower),
			"IsMark": reflect.ValueOf(unicode.IsMark),
			"IsNumber": reflect.ValueOf(unicode.IsNumber),
			"IsOneOf": reflect.ValueOf(unicode.IsOneOf),
			"IsPrint": reflect.ValueOf(unicode.IsPrint),
			"IsPunct": reflect.ValueOf(unicode.IsPunct),
			"IsSpace": reflect.ValueOf(unicode.IsSpace),
			"IsSymbol": reflect.ValueOf(unicode.IsSymbol),
			"IsTitle": reflect.ValueOf(unicode.IsTitle),
			"IsUpper": reflect.ValueOf(unicode.IsUpper),
			"SimpleFold": reflect.ValueOf(unicode.SimpleFold),
			"To": reflect.ValueOf(unicode.To),
			"ToLower": reflect.ValueOf(unicode.ToLower),
			"ToTitle": reflect.ValueOf(unicode.ToTitle),
			"ToUpper": reflect.ValueOf(unicode.ToUpper),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "atomic",
		Path:       "sync/atomic",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Value": reflect.TypeOf(func(atomic.Value){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"AddInt32": reflect.ValueOf(atomic.AddInt32),
			"AddInt64": reflect.ValueOf(atomic.AddInt64),
			"AddUint32": reflect.ValueOf(atomic.AddUint32),
			"AddUint64": reflect.ValueOf(atomic.AddUint64),
			"AddUintptr": reflect.ValueOf(atomic.AddUintptr),
			"CompareAndSwapInt32": reflect.ValueOf(atomic.CompareAndSwapInt32),
			"CompareAndSwapInt64": reflect.ValueOf(atomic.CompareAndSwapInt64),
			"CompareAndSwapPointer": reflect.ValueOf(atomic.CompareAndSwapPointer),
			"CompareAndSwapUint32": reflect.ValueOf(atomic.CompareAndSwapUint32),
			"CompareAndSwapUint64": reflect.ValueOf(atomic.CompareAndSwapUint64),
			"CompareAndSwapUintptr": reflect.ValueOf(atomic.CompareAndSwapUintptr),
			"LoadInt32": reflect.ValueOf(atomic.LoadInt32),
			"LoadInt64": reflect.ValueOf(atomic.LoadInt64),
			"LoadPointer": reflect.ValueOf(atomic.LoadPointer),
			"LoadUint32": reflect.ValueOf(atomic.LoadUint32),
			"LoadUint64": reflect.ValueOf(atomic.LoadUint64),
			"LoadUintptr": reflect.ValueOf(atomic.LoadUintptr),
			"StoreInt32": reflect.ValueOf(atomic.StoreInt32),
			"StoreInt64": reflect.ValueOf(atomic.StoreInt64),
			"StorePointer": reflect.ValueOf(atomic.StorePointer),
			"StoreUint32": reflect.ValueOf(atomic.StoreUint32),
			"StoreUint64": reflect.ValueOf(atomic.StoreUint64),
			"StoreUintptr": reflect.ValueOf(atomic.StoreUintptr),
			"SwapInt32": reflect.ValueOf(atomic.SwapInt32),
			"SwapInt64": reflect.ValueOf(atomic.SwapInt64),
			"SwapPointer": reflect.ValueOf(atomic.SwapPointer),
			"SwapUint32": reflect.ValueOf(atomic.SwapUint32),
			"SwapUint64": reflect.ValueOf(atomic.SwapUint64),
			"SwapUintptr": reflect.ValueOf(atomic.SwapUintptr),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "json",
		Path:       "encoding/json",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Decoder": reflect.TypeOf(func(json.Decoder){}).In(0),
			"Delim": reflect.TypeOf(func(json.Delim){}).In(0),
			"Encoder": reflect.TypeOf(func(json.Encoder){}).In(0),
			"InvalidUTF8Error": reflect.TypeOf(func(json.InvalidUTF8Error){}).In(0),
			"InvalidUnmarshalError": reflect.TypeOf(func(json.InvalidUnmarshalError){}).In(0),
			"Marshaler": reflect.TypeOf(func(json.Marshaler){}).In(0),
			"MarshalerError": reflect.TypeOf(func(json.MarshalerError){}).In(0),
			"Number": reflect.TypeOf(func(json.Number){}).In(0),
			"RawMessage": reflect.TypeOf(func(json.RawMessage){}).In(0),
			"SyntaxError": reflect.TypeOf(func(json.SyntaxError){}).In(0),
			"Token": reflect.TypeOf(func(json.Token){}).In(0),
			"UnmarshalFieldError": reflect.TypeOf(func(json.UnmarshalFieldError){}).In(0),
			"UnmarshalTypeError": reflect.TypeOf(func(json.UnmarshalTypeError){}).In(0),
			"Unmarshaler": reflect.TypeOf(func(json.Unmarshaler){}).In(0),
			"UnsupportedTypeError": reflect.TypeOf(func(json.UnsupportedTypeError){}).In(0),
			"UnsupportedValueError": reflect.TypeOf(func(json.UnsupportedValueError){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"Compact": reflect.ValueOf(json.Compact),
			"HTMLEscape": reflect.ValueOf(json.HTMLEscape),
			"Indent": reflect.ValueOf(json.Indent),
			"Marshal": reflect.ValueOf(json.Marshal),
			"MarshalIndent": reflect.ValueOf(json.MarshalIndent),
			"NewDecoder": reflect.ValueOf(json.NewDecoder),
			"NewEncoder": reflect.ValueOf(json.NewEncoder),
			"Unmarshal": reflect.ValueOf(json.Unmarshal),
			"Valid": reflect.ValueOf(json.Valid),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "template",
		Path:       "html/template",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CSS": reflect.TypeOf(func(template.CSS){}).In(0),
			"Error": reflect.TypeOf(func(template.Error){}).In(0),
			"ErrorCode": reflect.TypeOf(func(template.ErrorCode){}).In(0),
			"FuncMap": reflect.TypeOf(func(template.FuncMap){}).In(0),
			"HTML": reflect.TypeOf(func(template.HTML){}).In(0),
			"HTMLAttr": reflect.TypeOf(func(template.HTMLAttr){}).In(0),
			"JS": reflect.TypeOf(func(template.JS){}).In(0),
			"JSStr": reflect.TypeOf(func(template.JSStr){}).In(0),
			"Srcset": reflect.TypeOf(func(template.Srcset){}).In(0),
			"Template": reflect.TypeOf(func(template.Template){}).In(0),
			"URL": reflect.TypeOf(func(template.URL){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"HTMLEscape": reflect.ValueOf(template.HTMLEscape),
			"HTMLEscapeString": reflect.ValueOf(template.HTMLEscapeString),
			"HTMLEscaper": reflect.ValueOf(template.HTMLEscaper),
			"IsTrue": reflect.ValueOf(template.IsTrue),
			"JSEscape": reflect.ValueOf(template.JSEscape),
			"JSEscapeString": reflect.ValueOf(template.JSEscapeString),
			"JSEscaper": reflect.ValueOf(template.JSEscaper),
			"Must": reflect.ValueOf(template.Must),
			"New": reflect.ValueOf(template.New),
			"ParseFS": reflect.ValueOf(template.ParseFS),
			"ParseFiles": reflect.ValueOf(template.ParseFiles),
			"ParseGlob": reflect.ValueOf(template.ParseGlob),
			"URLQueryEscaper": reflect.ValueOf(template.URLQueryEscaper),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "path",
		Path:       "path",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"ErrBadPattern": reflect.ValueOf(func (error){}),

		},
		Funcs: map[string]reflect.Value{
			"Base": reflect.ValueOf(path.Base),
			"Clean": reflect.ValueOf(path.Clean),
			"Dir": reflect.ValueOf(path.Dir),
			"Ext": reflect.ValueOf(path.Ext),
			"IsAbs": reflect.ValueOf(path.IsAbs),
			"Join": reflect.ValueOf(path.Join),
			"Match": reflect.ValueOf(path.Match),
			"Split": reflect.ValueOf(path.Split),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "server",
		Path:       "github.com/linkxzhou/gongx/packages/server",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Context": reflect.TypeOf(func(server.Context){}).In(0),
			"HTTPError": reflect.TypeOf(func(server.HTTPError){}).In(0),
			"HTTPErrorHandler": reflect.TypeOf(func(server.HTTPErrorHandler){}).In(0),
			"HandlerFunc": reflect.TypeOf(func(server.HandlerFunc){}).In(0),
			"HttpServer": reflect.TypeOf(func(server.HttpServer){}).In(0),
			"Map": reflect.TypeOf(func(server.Map){}).In(0),
			"MiddlewareFunc": reflect.TypeOf(func(server.MiddlewareFunc){}).In(0),
			"Renderer": reflect.TypeOf(func(server.Renderer){}).In(0),
			"Response": reflect.TypeOf(func(server.Response){}).In(0),
			"Route": reflect.TypeOf(func(server.Route){}).In(0),
			"Router": reflect.TypeOf(func(server.Router){}).In(0),
			"Validator": reflect.TypeOf(func(server.Validator){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"ErrBadGateway": reflect.ValueOf(server.ErrBadGateway),
			"ErrBadRequest": reflect.ValueOf(server.ErrBadRequest),
			"ErrCookieNotFound": reflect.ValueOf(func (error){}),
			"ErrEncryptInvalid": reflect.ValueOf(func (error){}),
			"ErrForbidden": reflect.ValueOf(server.ErrForbidden),
			"ErrInternalServerError": reflect.ValueOf(server.ErrInternalServerError),
			"ErrInvalidCertOrKeyType": reflect.ValueOf(func (error){}),
			"ErrInvalidListenerNetwork": reflect.ValueOf(func (error){}),
			"ErrInvalidRedirectCode": reflect.ValueOf(func (error){}),
			"ErrMethodNotAllowed": reflect.ValueOf(server.ErrMethodNotAllowed),
			"ErrNotFound": reflect.ValueOf(server.ErrNotFound),
			"ErrRendererNotRegistered": reflect.ValueOf(func (error){}),
			"ErrRequestTimeout": reflect.ValueOf(server.ErrRequestTimeout),
			"ErrServiceUnavailable": reflect.ValueOf(server.ErrServiceUnavailable),
			"ErrSignInvalid": reflect.ValueOf(func (error){}),
			"ErrStatusRequestEntityTooLarge": reflect.ValueOf(server.ErrStatusRequestEntityTooLarge),
			"ErrTooManyRequests": reflect.ValueOf(server.ErrTooManyRequests),
			"ErrUnauthorized": reflect.ValueOf(server.ErrUnauthorized),
			"ErrUnsupportedMediaType": reflect.ValueOf(server.ErrUnsupportedMediaType),
			"ErrValidatorNotRegistered": reflect.ValueOf(func (error){}),
			"MethodNotAllowedHandler": reflect.ValueOf(server.MethodNotAllowedHandler),
			"NotFoundHandler": reflect.ValueOf(server.NotFoundHandler),

		},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(server.New),
			"NewHTTPError": reflect.ValueOf(server.NewHTTPError),
			"NewResponse": reflect.ValueOf(server.NewResponse),
			"NewRouter": reflect.ValueOf(server.NewRouter),
			"WrapHandler": reflect.ValueOf(server.WrapHandler),
			"WrapMiddleware": reflect.ValueOf(server.WrapMiddleware),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "multipart",
		Path:       "mime/multipart",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"File": reflect.TypeOf(func(multipart.File){}).In(0),
			"FileHeader": reflect.TypeOf(func(multipart.FileHeader){}).In(0),
			"Form": reflect.TypeOf(func(multipart.Form){}).In(0),
			"Part": reflect.TypeOf(func(multipart.Part){}).In(0),
			"Reader": reflect.TypeOf(func(multipart.Reader){}).In(0),
			"Writer": reflect.TypeOf(func(multipart.Writer){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"ErrMessageTooLarge": reflect.ValueOf(func (error){}),

		},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(multipart.NewReader),
			"NewWriter": reflect.ValueOf(multipart.NewWriter),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "tls",
		Path:       "crypto/tls",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Certificate": reflect.TypeOf(func(tls.Certificate){}).In(0),
			"CertificateRequestInfo": reflect.TypeOf(func(tls.CertificateRequestInfo){}).In(0),
			"CipherSuite": reflect.TypeOf(func(tls.CipherSuite){}).In(0),
			"ClientAuthType": reflect.TypeOf(func(tls.ClientAuthType){}).In(0),
			"ClientHelloInfo": reflect.TypeOf(func(tls.ClientHelloInfo){}).In(0),
			"ClientSessionCache": reflect.TypeOf(func(tls.ClientSessionCache){}).In(0),
			"ClientSessionState": reflect.TypeOf(func(tls.ClientSessionState){}).In(0),
			"Config": reflect.TypeOf(func(tls.Config){}).In(0),
			"Conn": reflect.TypeOf(func(tls.Conn){}).In(0),
			"ConnectionState": reflect.TypeOf(func(tls.ConnectionState){}).In(0),
			"CurveID": reflect.TypeOf(func(tls.CurveID){}).In(0),
			"Dialer": reflect.TypeOf(func(tls.Dialer){}).In(0),
			"RecordHeaderError": reflect.TypeOf(func(tls.RecordHeaderError){}).In(0),
			"RenegotiationSupport": reflect.TypeOf(func(tls.RenegotiationSupport){}).In(0),
			"SignatureScheme": reflect.TypeOf(func(tls.SignatureScheme){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"CipherSuiteName": reflect.ValueOf(tls.CipherSuiteName),
			"CipherSuites": reflect.ValueOf(tls.CipherSuites),
			"Client": reflect.ValueOf(tls.Client),
			"Dial": reflect.ValueOf(tls.Dial),
			"DialWithDialer": reflect.ValueOf(tls.DialWithDialer),
			"InsecureCipherSuites": reflect.ValueOf(tls.InsecureCipherSuites),
			"Listen": reflect.ValueOf(tls.Listen),
			"LoadX509KeyPair": reflect.ValueOf(tls.LoadX509KeyPair),
			"NewLRUClientSessionCache": reflect.ValueOf(tls.NewLRUClientSessionCache),
			"NewListener": reflect.ValueOf(tls.NewListener),
			"Server": reflect.ValueOf(tls.Server),
			"X509KeyPair": reflect.ValueOf(tls.X509KeyPair),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "sha1",
		Path:       "crypto/sha1",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(sha1.New),
			"Sum": reflect.ValueOf(sha1.Sum),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "heap",
		Path:       "container/heap",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Interface": reflect.TypeOf(func(heap.Interface){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"Fix": reflect.ValueOf(heap.Fix),
			"Init": reflect.ValueOf(heap.Init),
			"Pop": reflect.ValueOf(heap.Pop),
			"Push": reflect.ValueOf(heap.Push),
			"Remove": reflect.ValueOf(heap.Remove),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "list",
		Path:       "container/list",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Element": reflect.TypeOf(func(list.Element){}).In(0),
			"List": reflect.TypeOf(func(list.List){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(list.New),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "base64",
		Path:       "encoding/base64",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CorruptInputError": reflect.TypeOf(func(base64.CorruptInputError){}).In(0),
			"Encoding": reflect.TypeOf(func(base64.Encoding){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"RawStdEncoding": reflect.ValueOf(base64.RawStdEncoding),
			"RawURLEncoding": reflect.ValueOf(base64.RawURLEncoding),
			"StdEncoding": reflect.ValueOf(base64.StdEncoding),
			"URLEncoding": reflect.ValueOf(base64.URLEncoding),

		},
		Funcs: map[string]reflect.Value{
			"NewDecoder": reflect.ValueOf(base64.NewDecoder),
			"NewEncoder": reflect.ValueOf(base64.NewEncoder),
			"NewEncoding": reflect.ValueOf(base64.NewEncoding),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "hex",
		Path:       "encoding/hex",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"InvalidByteError": reflect.TypeOf(func(hex.InvalidByteError){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"ErrLength": reflect.ValueOf(func (error){}),

		},
		Funcs: map[string]reflect.Value{
			"Decode": reflect.ValueOf(hex.Decode),
			"DecodeString": reflect.ValueOf(hex.DecodeString),
			"DecodedLen": reflect.ValueOf(hex.DecodedLen),
			"Dump": reflect.ValueOf(hex.Dump),
			"Dumper": reflect.ValueOf(hex.Dumper),
			"Encode": reflect.ValueOf(hex.Encode),
			"EncodeToString": reflect.ValueOf(hex.EncodeToString),
			"EncodedLen": reflect.ValueOf(hex.EncodedLen),
			"NewDecoder": reflect.ValueOf(hex.NewDecoder),
			"NewEncoder": reflect.ValueOf(hex.NewEncoder),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "html",
		Path:       "html",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"EscapeString": reflect.ValueOf(html.EscapeString),
			"UnescapeString": reflect.ValueOf(html.UnescapeString),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "rand",
		Path:       "math/rand",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Rand": reflect.TypeOf(func(rand.Rand){}).In(0),
			"Source": reflect.TypeOf(func(rand.Source){}).In(0),
			"Source64": reflect.TypeOf(func(rand.Source64){}).In(0),
			"Zipf": reflect.TypeOf(func(rand.Zipf){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"ExpFloat64": reflect.ValueOf(rand.ExpFloat64),
			"Float32": reflect.ValueOf(rand.Float32),
			"Float64": reflect.ValueOf(rand.Float64),
			"Int": reflect.ValueOf(rand.Int),
			"Int31": reflect.ValueOf(rand.Int31),
			"Int31n": reflect.ValueOf(rand.Int31n),
			"Int63": reflect.ValueOf(rand.Int63),
			"Int63n": reflect.ValueOf(rand.Int63n),
			"Intn": reflect.ValueOf(rand.Intn),
			"New": reflect.ValueOf(rand.New),
			"NewSource": reflect.ValueOf(rand.NewSource),
			"NewZipf": reflect.ValueOf(rand.NewZipf),
			"NormFloat64": reflect.ValueOf(rand.NormFloat64),
			"Perm": reflect.ValueOf(rand.Perm),
			"Read": reflect.ValueOf(rand.Read),
			"Seed": reflect.ValueOf(rand.Seed),
			"Shuffle": reflect.ValueOf(rand.Shuffle),
			"Uint32": reflect.ValueOf(rand.Uint32),
			"Uint64": reflect.ValueOf(rand.Uint64),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "time",
		Path:       "time",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Duration": reflect.TypeOf(func(time.Duration){}).In(0),
			"Location": reflect.TypeOf(func(time.Location){}).In(0),
			"Month": reflect.TypeOf(func(time.Month){}).In(0),
			"ParseError": reflect.TypeOf(func(time.ParseError){}).In(0),
			"Ticker": reflect.TypeOf(func(time.Ticker){}).In(0),
			"Time": reflect.TypeOf(func(time.Time){}).In(0),
			"Timer": reflect.TypeOf(func(time.Timer){}).In(0),
			"Weekday": reflect.TypeOf(func(time.Weekday){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"Local": reflect.ValueOf(time.Local),
			"UTC": reflect.ValueOf(time.UTC),

		},
		Funcs: map[string]reflect.Value{
			"After": reflect.ValueOf(time.After),
			"AfterFunc": reflect.ValueOf(time.AfterFunc),
			"Date": reflect.ValueOf(time.Date),
			"FixedZone": reflect.ValueOf(time.FixedZone),
			"LoadLocation": reflect.ValueOf(time.LoadLocation),
			"LoadLocationFromTZData": reflect.ValueOf(time.LoadLocationFromTZData),
			"NewTicker": reflect.ValueOf(time.NewTicker),
			"NewTimer": reflect.ValueOf(time.NewTimer),
			"Now": reflect.ValueOf(time.Now),
			"Parse": reflect.ValueOf(time.Parse),
			"ParseDuration": reflect.ValueOf(time.ParseDuration),
			"ParseInLocation": reflect.ValueOf(time.ParseInLocation),
			"Since": reflect.ValueOf(time.Since),
			"Sleep": reflect.ValueOf(time.Sleep),
			"Tick": reflect.ValueOf(time.Tick),
			"Unix": reflect.ValueOf(time.Unix),
			"UnixMicro": reflect.ValueOf(time.UnixMicro),
			"UnixMilli": reflect.ValueOf(time.UnixMilli),
			"Until": reflect.ValueOf(time.Until),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "url",
		Path:       "net/url",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Error": reflect.TypeOf(func(url.Error){}).In(0),
			"EscapeError": reflect.TypeOf(func(url.EscapeError){}).In(0),
			"InvalidHostError": reflect.TypeOf(func(url.InvalidHostError){}).In(0),
			"URL": reflect.TypeOf(func(url.URL){}).In(0),
			"Userinfo": reflect.TypeOf(func(url.Userinfo){}).In(0),
			"Values": reflect.TypeOf(func(url.Values){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"Parse": reflect.ValueOf(url.Parse),
			"ParseQuery": reflect.ValueOf(url.ParseQuery),
			"ParseRequestURI": reflect.ValueOf(url.ParseRequestURI),
			"PathEscape": reflect.ValueOf(url.PathEscape),
			"PathUnescape": reflect.ValueOf(url.PathUnescape),
			"QueryEscape": reflect.ValueOf(url.QueryEscape),
			"QueryUnescape": reflect.ValueOf(url.QueryUnescape),
			"User": reflect.ValueOf(url.User),
			"UserPassword": reflect.ValueOf(url.UserPassword),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "io",
		Path:       "io",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"ByteReader": reflect.TypeOf(func(io.ByteReader){}).In(0),
			"ByteScanner": reflect.TypeOf(func(io.ByteScanner){}).In(0),
			"ByteWriter": reflect.TypeOf(func(io.ByteWriter){}).In(0),
			"Closer": reflect.TypeOf(func(io.Closer){}).In(0),
			"LimitedReader": reflect.TypeOf(func(io.LimitedReader){}).In(0),
			"PipeReader": reflect.TypeOf(func(io.PipeReader){}).In(0),
			"PipeWriter": reflect.TypeOf(func(io.PipeWriter){}).In(0),
			"ReadCloser": reflect.TypeOf(func(io.ReadCloser){}).In(0),
			"ReadSeekCloser": reflect.TypeOf(func(io.ReadSeekCloser){}).In(0),
			"ReadSeeker": reflect.TypeOf(func(io.ReadSeeker){}).In(0),
			"ReadWriteCloser": reflect.TypeOf(func(io.ReadWriteCloser){}).In(0),
			"ReadWriteSeeker": reflect.TypeOf(func(io.ReadWriteSeeker){}).In(0),
			"ReadWriter": reflect.TypeOf(func(io.ReadWriter){}).In(0),
			"Reader": reflect.TypeOf(func(io.Reader){}).In(0),
			"ReaderAt": reflect.TypeOf(func(io.ReaderAt){}).In(0),
			"ReaderFrom": reflect.TypeOf(func(io.ReaderFrom){}).In(0),
			"RuneReader": reflect.TypeOf(func(io.RuneReader){}).In(0),
			"RuneScanner": reflect.TypeOf(func(io.RuneScanner){}).In(0),
			"SectionReader": reflect.TypeOf(func(io.SectionReader){}).In(0),
			"Seeker": reflect.TypeOf(func(io.Seeker){}).In(0),
			"StringWriter": reflect.TypeOf(func(io.StringWriter){}).In(0),
			"WriteCloser": reflect.TypeOf(func(io.WriteCloser){}).In(0),
			"WriteSeeker": reflect.TypeOf(func(io.WriteSeeker){}).In(0),
			"Writer": reflect.TypeOf(func(io.Writer){}).In(0),
			"WriterAt": reflect.TypeOf(func(io.WriterAt){}).In(0),
			"WriterTo": reflect.TypeOf(func(io.WriterTo){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"Discard": reflect.ValueOf(func (io.Writer){}),
			"EOF": reflect.ValueOf(func (error){}),
			"ErrClosedPipe": reflect.ValueOf(func (error){}),
			"ErrNoProgress": reflect.ValueOf(func (error){}),
			"ErrShortBuffer": reflect.ValueOf(func (error){}),
			"ErrShortWrite": reflect.ValueOf(func (error){}),
			"ErrUnexpectedEOF": reflect.ValueOf(func (error){}),

		},
		Funcs: map[string]reflect.Value{
			"Copy": reflect.ValueOf(io.Copy),
			"CopyBuffer": reflect.ValueOf(io.CopyBuffer),
			"CopyN": reflect.ValueOf(io.CopyN),
			"LimitReader": reflect.ValueOf(io.LimitReader),
			"MultiReader": reflect.ValueOf(io.MultiReader),
			"MultiWriter": reflect.ValueOf(io.MultiWriter),
			"NewSectionReader": reflect.ValueOf(io.NewSectionReader),
			"NopCloser": reflect.ValueOf(io.NopCloser),
			"Pipe": reflect.ValueOf(io.Pipe),
			"ReadAll": reflect.ValueOf(io.ReadAll),
			"ReadAtLeast": reflect.ValueOf(io.ReadAtLeast),
			"ReadFull": reflect.ValueOf(io.ReadFull),
			"TeeReader": reflect.ValueOf(io.TeeReader),
			"WriteString": reflect.ValueOf(io.WriteString),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "ring",
		Path:       "container/ring",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Ring": reflect.TypeOf(func(ring.Ring){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(ring.New),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "sort",
		Path:       "sort",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Float64Slice": reflect.TypeOf(func(sort.Float64Slice){}).In(0),
			"IntSlice": reflect.TypeOf(func(sort.IntSlice){}).In(0),
			"Interface": reflect.TypeOf(func(sort.Interface){}).In(0),
			"StringSlice": reflect.TypeOf(func(sort.StringSlice){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"Float64s": reflect.ValueOf(sort.Float64s),
			"Float64sAreSorted": reflect.ValueOf(sort.Float64sAreSorted),
			"Ints": reflect.ValueOf(sort.Ints),
			"IntsAreSorted": reflect.ValueOf(sort.IntsAreSorted),
			"IsSorted": reflect.ValueOf(sort.IsSorted),
			"Reverse": reflect.ValueOf(sort.Reverse),
			"Search": reflect.ValueOf(sort.Search),
			"SearchFloat64s": reflect.ValueOf(sort.SearchFloat64s),
			"SearchInts": reflect.ValueOf(sort.SearchInts),
			"SearchStrings": reflect.ValueOf(sort.SearchStrings),
			"Slice": reflect.ValueOf(sort.Slice),
			"SliceIsSorted": reflect.ValueOf(sort.SliceIsSorted),
			"SliceStable": reflect.ValueOf(sort.SliceStable),
			"Sort": reflect.ValueOf(sort.Sort),
			"Stable": reflect.ValueOf(sort.Stable),
			"Strings": reflect.ValueOf(sort.Strings),
			"StringsAreSorted": reflect.ValueOf(sort.StringsAreSorted),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "fmt",
		Path:       "fmt",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Formatter": reflect.TypeOf(func(fmt.Formatter){}).In(0),
			"GoStringer": reflect.TypeOf(func(fmt.GoStringer){}).In(0),
			"ScanState": reflect.TypeOf(func(fmt.ScanState){}).In(0),
			"Scanner": reflect.TypeOf(func(fmt.Scanner){}).In(0),
			"State": reflect.TypeOf(func(fmt.State){}).In(0),
			"Stringer": reflect.TypeOf(func(fmt.Stringer){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"Errorf": reflect.ValueOf(fmt.Errorf),
			"Fprint": reflect.ValueOf(fmt.Fprint),
			"Fprintf": reflect.ValueOf(fmt.Fprintf),
			"Fprintln": reflect.ValueOf(fmt.Fprintln),
			"Fscan": reflect.ValueOf(fmt.Fscan),
			"Fscanf": reflect.ValueOf(fmt.Fscanf),
			"Fscanln": reflect.ValueOf(fmt.Fscanln),
			"Print": reflect.ValueOf(fmt.Print),
			"Printf": reflect.ValueOf(fmt.Printf),
			"Println": reflect.ValueOf(fmt.Println),
			"Scan": reflect.ValueOf(fmt.Scan),
			"Scanf": reflect.ValueOf(fmt.Scanf),
			"Scanln": reflect.ValueOf(fmt.Scanln),
			"Sprint": reflect.ValueOf(fmt.Sprint),
			"Sprintf": reflect.ValueOf(fmt.Sprintf),
			"Sprintln": reflect.ValueOf(fmt.Sprintln),
			"Sscan": reflect.ValueOf(fmt.Sscan),
			"Sscanf": reflect.ValueOf(fmt.Sscanf),
			"Sscanln": reflect.ValueOf(fmt.Sscanln),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "strconv",
		Path:       "strconv",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"NumError": reflect.TypeOf(func(strconv.NumError){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"ErrRange": reflect.ValueOf(func (error){}),
			"ErrSyntax": reflect.ValueOf(func (error){}),

		},
		Funcs: map[string]reflect.Value{
			"AppendBool": reflect.ValueOf(strconv.AppendBool),
			"AppendFloat": reflect.ValueOf(strconv.AppendFloat),
			"AppendInt": reflect.ValueOf(strconv.AppendInt),
			"AppendQuote": reflect.ValueOf(strconv.AppendQuote),
			"AppendQuoteRune": reflect.ValueOf(strconv.AppendQuoteRune),
			"AppendQuoteRuneToASCII": reflect.ValueOf(strconv.AppendQuoteRuneToASCII),
			"AppendQuoteRuneToGraphic": reflect.ValueOf(strconv.AppendQuoteRuneToGraphic),
			"AppendQuoteToASCII": reflect.ValueOf(strconv.AppendQuoteToASCII),
			"AppendQuoteToGraphic": reflect.ValueOf(strconv.AppendQuoteToGraphic),
			"AppendUint": reflect.ValueOf(strconv.AppendUint),
			"Atoi": reflect.ValueOf(strconv.Atoi),
			"CanBackquote": reflect.ValueOf(strconv.CanBackquote),
			"FormatBool": reflect.ValueOf(strconv.FormatBool),
			"FormatComplex": reflect.ValueOf(strconv.FormatComplex),
			"FormatFloat": reflect.ValueOf(strconv.FormatFloat),
			"FormatInt": reflect.ValueOf(strconv.FormatInt),
			"FormatUint": reflect.ValueOf(strconv.FormatUint),
			"IsGraphic": reflect.ValueOf(strconv.IsGraphic),
			"IsPrint": reflect.ValueOf(strconv.IsPrint),
			"Itoa": reflect.ValueOf(strconv.Itoa),
			"ParseBool": reflect.ValueOf(strconv.ParseBool),
			"ParseComplex": reflect.ValueOf(strconv.ParseComplex),
			"ParseFloat": reflect.ValueOf(strconv.ParseFloat),
			"ParseInt": reflect.ValueOf(strconv.ParseInt),
			"ParseUint": reflect.ValueOf(strconv.ParseUint),
			"Quote": reflect.ValueOf(strconv.Quote),
			"QuoteRune": reflect.ValueOf(strconv.QuoteRune),
			"QuoteRuneToASCII": reflect.ValueOf(strconv.QuoteRuneToASCII),
			"QuoteRuneToGraphic": reflect.ValueOf(strconv.QuoteRuneToGraphic),
			"QuoteToASCII": reflect.ValueOf(strconv.QuoteToASCII),
			"QuoteToGraphic": reflect.ValueOf(strconv.QuoteToGraphic),
			"QuotedPrefix": reflect.ValueOf(strconv.QuotedPrefix),
			"Unquote": reflect.ValueOf(strconv.Unquote),
			"UnquoteChar": reflect.ValueOf(strconv.UnquoteChar),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "utf8",
		Path:       "unicode/utf8",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"AppendRune": reflect.ValueOf(utf8.AppendRune),
			"DecodeLastRune": reflect.ValueOf(utf8.DecodeLastRune),
			"DecodeLastRuneInString": reflect.ValueOf(utf8.DecodeLastRuneInString),
			"DecodeRune": reflect.ValueOf(utf8.DecodeRune),
			"DecodeRuneInString": reflect.ValueOf(utf8.DecodeRuneInString),
			"EncodeRune": reflect.ValueOf(utf8.EncodeRune),
			"FullRune": reflect.ValueOf(utf8.FullRune),
			"FullRuneInString": reflect.ValueOf(utf8.FullRuneInString),
			"RuneCount": reflect.ValueOf(utf8.RuneCount),
			"RuneCountInString": reflect.ValueOf(utf8.RuneCountInString),
			"RuneLen": reflect.ValueOf(utf8.RuneLen),
			"RuneStart": reflect.ValueOf(utf8.RuneStart),
			"Valid": reflect.ValueOf(utf8.Valid),
			"ValidRune": reflect.ValueOf(utf8.ValidRune),
			"ValidString": reflect.ValueOf(utf8.ValidString),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "ioutil",
		Path:       "io/ioutil",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"Discard": reflect.ValueOf(func (io.Writer){}),

		},
		Funcs: map[string]reflect.Value{
			"NopCloser": reflect.ValueOf(ioutil.NopCloser),
			"ReadAll": reflect.ValueOf(ioutil.ReadAll),
			"ReadDir": reflect.ValueOf(ioutil.ReadDir),
			"ReadFile": reflect.ValueOf(ioutil.ReadFile),
			"TempDir": reflect.ValueOf(ioutil.TempDir),
			"TempFile": reflect.ValueOf(ioutil.TempFile),
			"WriteFile": reflect.ValueOf(ioutil.WriteFile),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "xml",
		Path:       "encoding/xml",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Attr": reflect.TypeOf(func(xml.Attr){}).In(0),
			"CharData": reflect.TypeOf(func(xml.CharData){}).In(0),
			"Comment": reflect.TypeOf(func(xml.Comment){}).In(0),
			"Decoder": reflect.TypeOf(func(xml.Decoder){}).In(0),
			"Directive": reflect.TypeOf(func(xml.Directive){}).In(0),
			"Encoder": reflect.TypeOf(func(xml.Encoder){}).In(0),
			"EndElement": reflect.TypeOf(func(xml.EndElement){}).In(0),
			"Marshaler": reflect.TypeOf(func(xml.Marshaler){}).In(0),
			"MarshalerAttr": reflect.TypeOf(func(xml.MarshalerAttr){}).In(0),
			"Name": reflect.TypeOf(func(xml.Name){}).In(0),
			"ProcInst": reflect.TypeOf(func(xml.ProcInst){}).In(0),
			"StartElement": reflect.TypeOf(func(xml.StartElement){}).In(0),
			"SyntaxError": reflect.TypeOf(func(xml.SyntaxError){}).In(0),
			"TagPathError": reflect.TypeOf(func(xml.TagPathError){}).In(0),
			"Token": reflect.TypeOf(func(xml.Token){}).In(0),
			"TokenReader": reflect.TypeOf(func(xml.TokenReader){}).In(0),
			"UnmarshalError": reflect.TypeOf(func(xml.UnmarshalError){}).In(0),
			"Unmarshaler": reflect.TypeOf(func(xml.Unmarshaler){}).In(0),
			"UnmarshalerAttr": reflect.TypeOf(func(xml.UnmarshalerAttr){}).In(0),
			"UnsupportedTypeError": reflect.TypeOf(func(xml.UnsupportedTypeError){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"HTMLAutoClose": reflect.ValueOf(xml.HTMLAutoClose),
			"HTMLEntity": reflect.ValueOf(xml.HTMLEntity),

		},
		Funcs: map[string]reflect.Value{
			"CopyToken": reflect.ValueOf(xml.CopyToken),
			"Escape": reflect.ValueOf(xml.Escape),
			"EscapeText": reflect.ValueOf(xml.EscapeText),
			"Marshal": reflect.ValueOf(xml.Marshal),
			"MarshalIndent": reflect.ValueOf(xml.MarshalIndent),
			"NewDecoder": reflect.ValueOf(xml.NewDecoder),
			"NewEncoder": reflect.ValueOf(xml.NewEncoder),
			"NewTokenDecoder": reflect.ValueOf(xml.NewTokenDecoder),
			"Unmarshal": reflect.ValueOf(xml.Unmarshal),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "math",
		Path:       "math",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"Abs": reflect.ValueOf(math.Abs),
			"Acos": reflect.ValueOf(math.Acos),
			"Acosh": reflect.ValueOf(math.Acosh),
			"Asin": reflect.ValueOf(math.Asin),
			"Asinh": reflect.ValueOf(math.Asinh),
			"Atan": reflect.ValueOf(math.Atan),
			"Atan2": reflect.ValueOf(math.Atan2),
			"Atanh": reflect.ValueOf(math.Atanh),
			"Cbrt": reflect.ValueOf(math.Cbrt),
			"Ceil": reflect.ValueOf(math.Ceil),
			"Copysign": reflect.ValueOf(math.Copysign),
			"Cos": reflect.ValueOf(math.Cos),
			"Cosh": reflect.ValueOf(math.Cosh),
			"Dim": reflect.ValueOf(math.Dim),
			"Erf": reflect.ValueOf(math.Erf),
			"Erfc": reflect.ValueOf(math.Erfc),
			"Erfcinv": reflect.ValueOf(math.Erfcinv),
			"Erfinv": reflect.ValueOf(math.Erfinv),
			"Exp": reflect.ValueOf(math.Exp),
			"Exp2": reflect.ValueOf(math.Exp2),
			"Expm1": reflect.ValueOf(math.Expm1),
			"FMA": reflect.ValueOf(math.FMA),
			"Float32bits": reflect.ValueOf(math.Float32bits),
			"Float32frombits": reflect.ValueOf(math.Float32frombits),
			"Float64bits": reflect.ValueOf(math.Float64bits),
			"Float64frombits": reflect.ValueOf(math.Float64frombits),
			"Floor": reflect.ValueOf(math.Floor),
			"Frexp": reflect.ValueOf(math.Frexp),
			"Gamma": reflect.ValueOf(math.Gamma),
			"Hypot": reflect.ValueOf(math.Hypot),
			"Ilogb": reflect.ValueOf(math.Ilogb),
			"Inf": reflect.ValueOf(math.Inf),
			"IsInf": reflect.ValueOf(math.IsInf),
			"IsNaN": reflect.ValueOf(math.IsNaN),
			"J0": reflect.ValueOf(math.J0),
			"J1": reflect.ValueOf(math.J1),
			"Jn": reflect.ValueOf(math.Jn),
			"Ldexp": reflect.ValueOf(math.Ldexp),
			"Lgamma": reflect.ValueOf(math.Lgamma),
			"Log": reflect.ValueOf(math.Log),
			"Log10": reflect.ValueOf(math.Log10),
			"Log1p": reflect.ValueOf(math.Log1p),
			"Log2": reflect.ValueOf(math.Log2),
			"Logb": reflect.ValueOf(math.Logb),
			"Max": reflect.ValueOf(math.Max),
			"Min": reflect.ValueOf(math.Min),
			"Mod": reflect.ValueOf(math.Mod),
			"Modf": reflect.ValueOf(math.Modf),
			"NaN": reflect.ValueOf(math.NaN),
			"Nextafter": reflect.ValueOf(math.Nextafter),
			"Nextafter32": reflect.ValueOf(math.Nextafter32),
			"Pow": reflect.ValueOf(math.Pow),
			"Pow10": reflect.ValueOf(math.Pow10),
			"Remainder": reflect.ValueOf(math.Remainder),
			"Round": reflect.ValueOf(math.Round),
			"RoundToEven": reflect.ValueOf(math.RoundToEven),
			"Signbit": reflect.ValueOf(math.Signbit),
			"Sin": reflect.ValueOf(math.Sin),
			"Sincos": reflect.ValueOf(math.Sincos),
			"Sinh": reflect.ValueOf(math.Sinh),
			"Sqrt": reflect.ValueOf(math.Sqrt),
			"Tan": reflect.ValueOf(math.Tan),
			"Tanh": reflect.ValueOf(math.Tanh),
			"Trunc": reflect.ValueOf(math.Trunc),
			"Y0": reflect.ValueOf(math.Y0),
			"Y1": reflect.ValueOf(math.Y1),
			"Yn": reflect.ValueOf(math.Yn),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "sync",
		Path:       "sync",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Cond": reflect.TypeOf(func(sync.Cond){}).In(0),
			"Locker": reflect.TypeOf(func(sync.Locker){}).In(0),
			"Map": reflect.TypeOf(func(sync.Map){}).In(0),
			"Mutex": reflect.TypeOf(func(sync.Mutex){}).In(0),
			"Once": reflect.TypeOf(func(sync.Once){}).In(0),
			"Pool": reflect.TypeOf(func(sync.Pool){}).In(0),
			"RWMutex": reflect.TypeOf(func(sync.RWMutex){}).In(0),
			"WaitGroup": reflect.TypeOf(func(sync.WaitGroup){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{

		},
		Funcs: map[string]reflect.Value{
			"NewCond": reflect.ValueOf(sync.NewCond),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
	RegisterPackage(&Package{
		Name:       "binary",
		Path:       "encoding/binary",
		Deps:       make(map[string]string),
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"ByteOrder": reflect.TypeOf(func(binary.ByteOrder){}).In(0),

		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{
			"BigEndian": reflect.ValueOf(binary.BigEndian),
			"LittleEndian": reflect.ValueOf(binary.LittleEndian),

		},
		Funcs: map[string]reflect.Value{
			"PutUvarint": reflect.ValueOf(binary.PutUvarint),
			"PutVarint": reflect.ValueOf(binary.PutVarint),
			"Read": reflect.ValueOf(binary.Read),
			"ReadUvarint": reflect.ValueOf(binary.ReadUvarint),
			"ReadVarint": reflect.ValueOf(binary.ReadVarint),
			"Size": reflect.ValueOf(binary.Size),
			"Uvarint": reflect.ValueOf(binary.Uvarint),
			"Varint": reflect.ValueOf(binary.Varint),
			"Write": reflect.ValueOf(binary.Write),

		},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	})
		
}
	
