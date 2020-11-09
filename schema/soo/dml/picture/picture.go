//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package picture ;import (_b "encoding/xml";_gf "github.com/unidoc/unioffice";_ga "github.com/unidoc/unioffice/schema/soo/dml";);func NewCT_PictureNonVisual ()*CT_PictureNonVisual {_de :=&CT_PictureNonVisual {};_de .CNvPr =_ga .NewCT_NonVisualDrawingProps ();_de .CNvPicPr =_ga .NewCT_NonVisualPictureProperties ();return _de ;};func (_aa *CT_PictureNonVisual )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_aa .CNvPr =_ga .NewCT_NonVisualDrawingProps ();_aa .CNvPicPr =_ga .NewCT_NonVisualPictureProperties ();_baf :for {_ac ,_fab :=d .Token ();if _fab !=nil {return _fab ;};switch _ea :=_ac .(type ){case _b .StartElement :switch _ea .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0063\u004e\u0076P\u0072"},_b .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0063\u004e\u0076P\u0072"}:if _ae :=d .DecodeElement (_aa .CNvPr ,&_ea );_ae !=nil {return _ae ;};case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"},_b .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}:if _ad :=d .DecodeElement (_aa .CNvPicPr ,&_ea );_ad !=nil {return _ad ;};default:_gf .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065No\u006e\u0056\u0069\u0073\u0075\u0061\u006c\u0020\u0025\u0076",_ea .Name );if _bad :=d .Skip ();_bad !=nil {return _bad ;};};case _b .EndElement :break _baf ;case _b .CharData :};};return nil ;};func (_aga *Pic )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_aga .CT_Picture =*NewCT_Picture ();_fcg :for {_be ,_dea :=d .Token ();if _dea !=nil {return _dea ;};switch _afg :=_be .(type ){case _b .StartElement :switch _afg .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u006ev\u0050\u0069\u0063\u0050\u0072"},_b .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u006ev\u0050\u0069\u0063\u0050\u0072"}:if _ca :=d .DecodeElement (_aga .NvPicPr ,&_afg );_ca !=nil {return _ca ;};case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"},_b .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:if _fgg :=d .DecodeElement (_aga .BlipFill ,&_afg );_fgg !=nil {return _fgg ;};case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0073\u0070\u0050\u0072"},_b .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0073\u0070\u0050\u0072"}:if _cb :=d .DecodeElement (_aga .SpPr ,&_afg );_cb !=nil {return _cb ;};default:_gf .Log ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006fn\u0020\u0050i\u0063\u0020\u0025\u0076",_afg .Name );if _ebe :=d .Skip ();_ebe !=nil {return _ebe ;};};case _b .EndElement :break _fcg ;case _b .CharData :};};return nil ;};func (_bbc *Pic )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"});start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0070\u0069c"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065"});start .Attr =append (start .Attr ,_b .Attr {Name :_b .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0070i\u0063\u003a\u0070\u0069\u0063";return _bbc .CT_Picture .MarshalXML (e ,start );};

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (_cc *CT_PictureNonVisual )ValidateWithPath (path string )error {if _acb :=_cc .CNvPr .ValidateWithPath (path +"\u002f\u0043\u004e\u0076\u0050\u0072");_acb !=nil {return _acb ;};if _ab :=_cc .CNvPicPr .ValidateWithPath (path +"\u002fC\u004e\u0076\u0050\u0069\u0063\u0050r");_ab !=nil {return _ab ;};return nil ;};func (_bac *CT_PictureNonVisual )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {e .EncodeToken (start );_ge :=_b .StartElement {Name :_b .Name {Local :"\u0070i\u0063\u003a\u0063\u004e\u0076\u0050r"}};e .EncodeElement (_bac .CNvPr ,_ge );_af :=_b .StartElement {Name :_b .Name {Local :"\u0070\u0069\u0063:\u0063\u004e\u0076\u0050\u0069\u0063\u0050\u0072"}};e .EncodeElement (_bac .CNvPicPr ,_af );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};func (_f *CT_Picture )MarshalXML (e *_b .Encoder ,start _b .StartElement )error {e .EncodeToken (start );_a :=_b .StartElement {Name :_b .Name {Local :"p\u0069\u0063\u003a\u006e\u0076\u0050\u0069\u0063\u0050\u0072"}};e .EncodeElement (_f .NvPicPr ,_a );_bb :=_b .StartElement {Name :_b .Name {Local :"\u0070\u0069\u0063:\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}};e .EncodeElement (_f .BlipFill ,_bb );_d :=_b .StartElement {Name :_b .Name {Local :"\u0070\u0069\u0063\u003a\u0073\u0070\u0050\u0072"}};e .EncodeElement (_f .SpPr ,_d );e .EncodeToken (_b .EndElement {Name :start .Name });return nil ;};type CT_PictureNonVisual struct{CNvPr *_ga .CT_NonVisualDrawingProps ;CNvPicPr *_ga .CT_NonVisualPictureProperties ;};type Pic struct{CT_Picture };func (_dc *CT_Picture )UnmarshalXML (d *_b .Decoder ,start _b .StartElement )error {_dc .NvPicPr =NewCT_PictureNonVisual ();_dc .BlipFill =_ga .NewCT_BlipFillProperties ();_dc .SpPr =_ga .NewCT_ShapeProperties ();_ag :for {_ba ,_da :=d .Token ();if _da !=nil {return _da ;};switch _fg :=_ba .(type ){case _b .StartElement :switch _fg .Name {case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u006ev\u0050\u0069\u0063\u0050\u0072"},_b .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u006ev\u0050\u0069\u0063\u0050\u0072"}:if _fa :=d .DecodeElement (_dc .NvPicPr ,&_fg );_fa !=nil {return _fa ;};case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"},_b .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0062\u006c\u0069\u0070\u0046\u0069\u006c\u006c"}:if _e :=d .DecodeElement (_dc .BlipFill ,&_fg );_e !=nil {return _e ;};case _b .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065",Local :"\u0073\u0070\u0050\u0072"},_b .Name {Space :"\u0068t\u0074\u0070:\u002f\u002f\u0070\u0075r\u006c\u002e\u006fc\u006c\u0063\u002e\u006f\u0072\u0067\u002f\u006f\u006fxm\u006c\u002f\u0064r\u0061\u0077i\u006e\u0067\u006d\u006c\u002f\u0070i\u0063\u0074u\u0072\u0065",Local :"\u0073\u0070\u0050\u0072"}:if _ec :=d .DecodeElement (_dc .SpPr ,&_fg );_ec !=nil {return _ec ;};default:_gf .Log ("\u0073k\u0069\u0070p\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006ce\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005fP\u0069\u0063\u0074\u0075\u0072\u0065\u0020\u0025\u0076",_fg .Name );if _fc :=d .Skip ();_fc !=nil {return _fc ;};};case _b .EndElement :break _ag ;case _b .CharData :};};return nil ;};

// Validate validates the CT_Picture and its children
func (_bc *CT_Picture )Validate ()error {return _bc .ValidateWithPath ("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065");};

// Validate validates the CT_PictureNonVisual and its children
func (_cf *CT_PictureNonVisual )Validate ()error {return _cf .ValidateWithPath ("\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c");};type CT_Picture struct{NvPicPr *CT_PictureNonVisual ;BlipFill *_ga .CT_BlipFillProperties ;SpPr *_ga .CT_ShapeProperties ;};func NewPic ()*Pic {_ebb :=&Pic {};_ebb .CT_Picture =*NewCT_Picture ();return _ebb };

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_dac *CT_Picture )ValidateWithPath (path string )error {if _eb :=_dac .NvPicPr .ValidateWithPath (path +"\u002f\u004e\u0076\u0050\u0069\u0063\u0050\u0072");_eb !=nil {return _eb ;};if _fb :=_dac .BlipFill .ValidateWithPath (path +"\u002fB\u006c\u0069\u0070\u0046\u0069\u006cl");_fb !=nil {return _fb ;};if _ef :=_dac .SpPr .ValidateWithPath (path +"\u002f\u0053\u0070P\u0072");_ef !=nil {return _ef ;};return nil ;};func NewCT_Picture ()*CT_Picture {_c :=&CT_Picture {};_c .NvPicPr =NewCT_PictureNonVisual ();_c .BlipFill =_ga .NewCT_BlipFillProperties ();_c .SpPr =_ga .NewCT_ShapeProperties ();return _c ;};

// ValidateWithPath validates the Pic and its children, prefixing error messages with path
func (_fbd *Pic )ValidateWithPath (path string )error {if _gd :=_fbd .CT_Picture .ValidateWithPath (path );_gd !=nil {return _gd ;};return nil ;};

// Validate validates the Pic and its children
func (_gfd *Pic )Validate ()error {return _gfd .ValidateWithPath ("\u0050\u0069\u0063")};func init (){_gf .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065","\u0043\u0054\u005f\u0050ic\u0074\u0075\u0072\u0065\u004e\u006f\u006e\u0056\u0069\u0073\u0075\u0061\u006c",NewCT_PictureNonVisual );_gf .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065","\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065",NewCT_Picture );_gf .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068e\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006frg\u002f\u0064\u0072\u0061\u0077\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u0070\u0069\u0063\u0074\u0075\u0072\u0065","\u0070\u0069\u0063",NewPic );};