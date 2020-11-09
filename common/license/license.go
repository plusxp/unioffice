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

// Package license helps manage commercial licenses and check if they are valid for the version of unidoc used.
package license ;import (_g "bytes";_bd "compress/gzip";_bb "crypto";_ae "crypto/rand";_deb "crypto/rsa";_f "crypto/sha256";_dd "crypto/sha512";_fe "crypto/x509";_bf "encoding/base64";_ag "encoding/hex";_de "encoding/json";_cb "encoding/pem";_ba "errors";_e "fmt";_fg "github.com/unidoc/unioffice/common";_c "io";_d "log";_bc "regexp";_ce "strings";_a "time";);func _bg (_bbf string ,_ca string )([]byte ,error ){_fb :="\u000a\u002b\u000a";_dc :="\u000d\u000a\u002b\r\u000a";_da :=_ce .Index (_ca ,_fb );if _da ==-1{_da =_ce .Index (_ca ,_dc );if _da ==-1{return nil ,_e .Errorf ("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u002c \u0073i\u0067n\u0061t\u0075\u0072\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u006f\u0072");};};_aec :=_ca [:_da ];_ddc :=_da +len (_fb );_ddg :=_ca [_ddc :];if _aec ==""||_ddg ==""{return nil ,_e .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0069n\u0070\u0075\u0074,\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0069\u0067\u0069n\u0061\u006c \u006f\u0072\u0020\u0073\u0069\u0067n\u0061\u0074u\u0072\u0065");};_gf ,_deg :=_bf .StdEncoding .DecodeString (_aec );if _deg !=nil {return nil ,_e .Errorf ("\u0069\u006e\u0076\u0061li\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0072\u0069\u0067\u0069\u006ea\u006c");};_eb ,_deg :=_bf .StdEncoding .DecodeString (_ddg );if _deg !=nil {return nil ,_e .Errorf ("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");};_aga ,_ :=_cb .Decode ([]byte (_bbf ));if _aga ==nil {return nil ,_e .Errorf ("\u0050\u0075\u0062\u004b\u0065\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064");};_fa ,_deg :=_fe .ParsePKIXPublicKey (_aga .Bytes );if _deg !=nil {return nil ,_deg ;};_bdb :=_fa .(*_deb .PublicKey );if _bdb ==nil {return nil ,_e .Errorf ("\u0050u\u0062\u004b\u0065\u0079\u0020\u0063\u006f\u006e\u0076\u0065\u0072s\u0069\u006f\u006e\u0020\u0066\u0061\u0069\u006c\u0065\u0064");};_fad :=_dd .New ();_fad .Write (_gf );_ced :=_fad .Sum (nil );_deg =_deb .VerifyPKCS1v15 (_bdb ,_bb .SHA512 ,_ced ,_eb );if _deg !=nil {return nil ,_deg ;};return _gf ,nil ;};func (_eec *LicenseKey )isExpired ()bool {return _eec .getExpiryDateToCompare ().After (_eec .ExpiresAt )};var _faga =MakeUnlicensedKey ();const _aeg ="\u0033\u0030\u0035\u0063\u0033\u0030\u0030\u00640\u0036\u0030\u0039\u0032\u0061\u0038\u00364\u0038\u0038\u0036\u0066\u0037\u0030d\u0030\u0031\u0030\u0031\u0030\u00310\u0035\u0030\u0030\u0030\u0033\u0034\u0062\u0030\u0030\u0033\u0030\u00348\u0030\u0032\u0034\u0031\u0030\u0030\u0062\u0038\u0037\u0065\u0061\u0066\u0062\u0036\u0063\u0030\u0037\u0034\u0039\u0039\u0065\u0062\u00397\u0063\u0063\u0039\u0064\u0033\u0035\u0036\u0035\u0065\u0063\u00663\u0031\u0036\u0038\u0031\u0039\u0036\u0033\u0030\u0031\u0039\u0030\u0037c\u0038\u0034\u0031\u0061\u0064\u0064c6\u0036\u0035\u0030\u0038\u0036\u0062\u0062\u0033\u0065\u0064\u0038\u0065\u0062\u0031\u0032\u0064\u0039\u0064\u0061\u0032\u0036\u0063\u0061\u0066\u0061\u0039\u0036\u00345\u0030\u00314\u0036\u0064\u0061\u0038\u0062\u0064\u0030\u0063c\u0066\u0031\u0035\u0035\u0066\u0063a\u0063\u0063\u00368\u0036\u0039\u0035\u0035\u0065\u0066\u0030\u0033\u0030\u0032\u0066\u0061\u0034\u0034\u0061\u0061\u0033\u0065\u0063\u0038\u0039\u0034\u0031\u0037\u0062\u0030\u0032\u0030\u0033\u0030\u0031\u0030\u0030\u0030\u0031";func GetLicenseKey ()*LicenseKey {if _faga ==nil {return nil ;};_gbc :=*_faga ;return &_gbc ;};

// MakeUnlicensedKey returns an unlicensed key.
func MakeUnlicensedKey ()*LicenseKey {_agg :=LicenseKey {};_agg .CustomerName ="\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";_agg .Tier =LicenseTierUnlicensed ;_agg .CreatedAt =_a .Now ().UTC ();_agg .CreatedAtInt =_agg .CreatedAt .Unix ();return &_agg ;};func _ad (_ga string ,_gd []byte )(string ,error ){_ddd ,_ :=_cb .Decode ([]byte (_ga ));if _ddd ==nil {return "",_e .Errorf ("\u0050\u0072\u0069\u0076\u004b\u0065\u0079\u0020\u0066a\u0069\u006c\u0065\u0064");};_adb ,_ee :=_fe .ParsePKCS1PrivateKey (_ddd .Bytes );if _ee !=nil {return "",_ee ;};_bcf :=_dd .New ();_bcf .Write (_gd );_gdb :=_bcf .Sum (nil );_eed ,_ee :=_deb .SignPKCS1v15 (_ae .Reader ,_adb ,_bb .SHA512 ,_gdb );if _ee !=nil {return "",_ee ;};_cc :=_bf .StdEncoding .EncodeToString (_gd );_cc +="\u000a\u002b\u000a";_cc +=_bf .StdEncoding .EncodeToString (_eed );return _cc ,nil ;};const (_ab ="\u002d\u002d\u002d--\u0042\u0045\u0047\u0049\u004e\u0020\u0055\u004e\u0049D\u004fC\u0020L\u0049C\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d";_df ="\u002d\u002d\u002d\u002d\u002d\u0045\u004e\u0044\u0020\u0055\u004e\u0049\u0044\u004f\u0043 \u004cI\u0043\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d";);

// Verify verifies a license by checking the license content and signature
// against a public key.
func (_gbd LegacyLicense )Verify (pubKey *_deb .PublicKey )error {_fd :=_gbd ;_fd .Signature ="";_bda :=_g .Buffer {};_gdfb :=_de .NewEncoder (&_bda );if _dea :=_gdfb .Encode (_fd );_dea !=nil {return _dea ;};_cde ,_bba :=_ag .DecodeString (_gbd .Signature );if _bba !=nil {return _bba ;};_dbf :=_f .Sum256 (_bda .Bytes ());_bba =_deb .VerifyPKCS1v15 (pubKey ,_bb .SHA256 ,_dbf [:],_cde );return _bba ;};var _adc =_a .Date (2010,1,1,0,0,0,0,_a .UTC );

// TypeToString returns a string representation of the license type.
func (_bdg *LicenseKey )TypeToString ()string {if _bdg .Tier ==LicenseTierUnlicensed {return "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";};if _bdg .Tier ==LicenseTierCommunity {return "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";};if _bdg .Tier ==LicenseTierIndividual ||_bdg .Tier =="\u0069\u006e\u0064i\u0065"{return "\u0043\u006f\u006dm\u0065\u0072\u0063\u0069a\u006c\u0020\u004c\u0069\u0063\u0065\u006es\u0065\u0020\u002d\u0020\u0049\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c";};return "\u0043\u006fm\u006d\u0065\u0072\u0063\u0069\u0061\u006c\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u002d\u0020\u0042\u0075\u0073\u0069ne\u0073\u0073";};func _db (_gb string )(LicenseKey ,error ){var _ebb LicenseKey ;_agc ,_ebf :=_baf (_ab ,_df ,_gb );if _ebf !=nil {return _ebb ,_ebf ;};_agd ,_ebf :=_bg (_bgf ,_agc );if _ebf !=nil {return _ebb ,_ebf ;};_ebf =_de .Unmarshal (_agd ,&_ebb );if _ebf !=nil {return _ebb ,_ebf ;};_ebb .CreatedAt =_a .Unix (_ebb .CreatedAtInt ,0);_ebb .ExpiresAt =_a .Unix (_ebb .ExpiresAtInt ,0);return _ebb ,nil ;};

// LegacyLicenseType is the type of license
type LegacyLicenseType byte ;

// ToString returns a string representing the license information.
func (_gg *LicenseKey )ToString ()string {_gff :=_e .Sprintf ("\u004ci\u0063e\u006e\u0073\u0065\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a",_gg .LicenseId );_gff +=_e .Sprintf ("\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a",_gg .CustomerId );_gff +=_e .Sprintf ("\u0043u\u0073t\u006f\u006d\u0065\u0072\u0020N\u0061\u006de\u003a\u0020\u0025\u0073\u000a",_gg .CustomerName );_gff +=_e .Sprintf ("\u0054i\u0065\u0072\u003a\u0020\u0025\u0073\n",_gg .Tier );_gff +=_e .Sprintf ("\u0043r\u0065a\u0074\u0065\u0064\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a",_fg .UtcTimeFormat (_gg .CreatedAt ));_gff +=_e .Sprintf ("\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a",_fg .UtcTimeFormat (_gg .ExpiresAt ));_gff +=_e .Sprintf ("\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u003a\u0020\u0025\u0073\u0020<\u0025\u0073\u003e\u000a",_gg .CreatorName ,_gg .CreatorEmail );return _gff ;};var _aa =false ;func _baf (_dfc string ,_ea string ,_cd string )(string ,error ){_ef :=_ce .Index (_cd ,_dfc );if _ef ==-1{return "",_e .Errorf ("\u0068\u0065a\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_fag :=_ce .Index (_cd ,_ea );if _fag ==-1{return "",_e .Errorf ("\u0066\u006fo\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_cf :=_ef +len (_dfc )+1;return _cd [_cf :_fag -1],nil ;};var _ggc *_deb .PublicKey ;

// Validate returns an error if the licenseis invalid, nil otherwise.
func (_gc *LicenseKey )Validate ()error {if len (_gc .LicenseId )< 10{return _e .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020L\u0069\u0063\u0065n\u0073e\u0020\u0049\u0064");};if len (_gc .CustomerId )< 10{return _e .Errorf ("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020C\u0075\u0073\u0074\u006f\u006d\u0065\u0072 \u0049\u0064");};if len (_gc .CustomerName )< 1{return _e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043u\u0073\u0074\u006f\u006d\u0065\u0072\u0020\u004e\u0061\u006d\u0065");};if _adc .After (_gc .CreatedAt ){return _e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064 \u0041\u0074\u0020\u0069\u0073 \u0069\u006ev\u0061\u006c\u0069\u0064");};if _gc .CreatedAt .After (_gc .ExpiresAt ){return _e .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0041\u0074 \u0063a\u006e\u006e\u006f\u0074 \u0062\u0065 \u0047\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0045\u0078\u0070\u0069\u0072\u0065\u0073\u0020\u0041\u0074");};if _gc .isExpired (){return _e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020l\u0069\u0063\u0065ns\u0065\u003a\u0020\u0054\u0068\u0065 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0065\u0078\u0070\u0069r\u0065\u0064");};if len (_gc .CreatorName )< 1{return _e .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020na\u006d\u0065");};if len (_gc .CreatorEmail )< 1{return _e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043r\u0065\u0061\u0074\u006f\u0072\u0020\u0065\u006d\u0061\u0069\u006c");};if !_gc .UniOffice {return _e .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0054\u0068\u0069s\u0020\u0055\u006e\u0069\u0044\u006f\u0063\u0020\u006be\u0079\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020f\u006f\u0072\u0020\u0055\u006e\u0069O\u0066\u0066\u0069c\u0065\u002e");};return nil ;};func init (){_ac ,_bafb :=_ag .DecodeString (_aeg );if _bafb !=nil {_d .Fatalf ("e\u0072\u0072\u006f\u0072 r\u0065a\u0064\u0069\u006e\u0067\u0020k\u0065\u0079\u003a\u0020\u0025\u0073",_bafb );};_abb ,_bafb :=_fe .ParsePKIXPublicKey (_ac );if _bafb !=nil {_d .Fatalf ("e\u0072\u0072\u006f\u0072 r\u0065a\u0064\u0069\u006e\u0067\u0020k\u0065\u0079\u003a\u0020\u0025\u0073",_bafb );};_ggc =_abb .(*_deb .PublicKey );};func (_gdf *LicenseKey )getExpiryDateToCompare ()_a .Time {if _gdf .Trial {return _a .Now ().UTC ();};return _fg .ReleasedAt ;};

// IsLicensed returns true if the package is licensed.
func (_bge *LicenseKey )IsLicensed ()bool {return _bge .Tier !=LicenseTierUnlicensed };

// SetLicenseKey sets and validates the license key.
func SetLicenseKey (content string ,customerName string )error {if _aa {return nil ;};_aag ,_acg :=_db (content );if _acg !=nil {return _acg ;};if _ce .ToLower (_aag .CustomerName )!=_ce .ToLower (customerName ){return _e .Errorf ("\u0063\u0075\u0073\u0074\u006fm\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006d\u0061t\u0063\u0068\u002c\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067\u006f\u0074\u0020\u0027\u0025\u0073'",customerName ,_aag .CustomerName );};_acg =_aag .Validate ();if _acg !=nil {return _acg ;};_faga =&_aag ;return nil ;};

// LegacyLicense holds the old-style unioffice license information.
type LegacyLicense struct{Name string ;Signature string `json:",omitempty"`;Expiration _a .Time ;LicenseType LegacyLicenseType ;};

// SetLegacyLicenseKey installs a legacy license code. License codes issued prior to June 2019.
// Will be removed at some point in a future major version.
func SetLegacyLicenseKey (s string )error {_eaf :=_bc .MustCompile ("\u005c\u0073");s =_eaf .ReplaceAllString (s ,"");var _efd _c .Reader ;_efd =_ce .NewReader (s );_efd =_bf .NewDecoder (_bf .RawURLEncoding ,_efd );_efd ,_fga :=_bd .NewReader (_efd );if _fga !=nil {return _fga ;};_fbb :=_de .NewDecoder (_efd );_abd :=&LegacyLicense {};if _bfc :=_fbb .Decode (_abd );_bfc !=nil {return _bfc ;};if _eafa :=_abd .Verify (_ggc );_eafa !=nil {return _ba .New ("\u006c\u0069\u0063en\u0073\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006e\u0020\u0065\u0072\u0072\u006f\u0072");};if _abd .Expiration .Before (_fg .ReleasedAt ){return _ba .New ("\u006ci\u0063e\u006e\u0073\u0065\u0020\u0065\u0078\u0070\u0069\u0072\u0065\u0064");};_eg :=_a .Now ().UTC ();_dg :=LicenseKey {};_dg .CreatedAt =_eg ;_dg .CustomerId ="\u004c\u0065\u0067\u0061\u0063\u0079";_dg .CustomerName =_abd .Name ;_dg .Tier =LicenseTierBusiness ;_dg .ExpiresAt =_abd .Expiration ;_dg .CreatorName ="\u0055\u006e\u0069\u0044\u006f\u0063\u0020\u0073\u0075p\u0070\u006f\u0072\u0074";_dg .CreatorEmail ="\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0040\u0075\u006e\u0069\u0064o\u0063\u002e\u0069\u006f";_dg .UniOffice =true ;_faga =&_dg ;return nil ;};const _bgf ="\u000a\u002d\u002d\u002d\u002d\u002d\u0042\u0045\u0047\u0049\u004e \u0050\u0055\u0042\u004c\u0049\u0043\u0020\u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\u000a\u004d\u0049I\u0042\u0049\u006a\u0041NB\u0067\u006b\u0071\u0068\u006b\u0069G\u0039\u0077\u0030\u0042\u0041\u0051\u0045\u0046A\u0041\u004f\u0043\u0041\u0051\u0038\u0041\u004d\u0049\u0049\u0042\u0043\u0067\u004b\u0043\u0041\u0051\u0045A\u006dF\u0055\u0069\u0079\u0064\u0037\u0062\u0035\u0058\u006a\u0070\u006b\u0050\u0035\u0052\u0061\u0070\u0034\u0077\u000a\u0044\u0063\u0031d\u0079\u007a\u0049\u0051\u0034\u004c\u0065\u006b\u0078\u0072\u0076\u0079\u0074\u006e\u0045\u004d\u0070\u004e\u0055\u0062\u006f\u0036i\u0041\u0037\u0034\u0056\u0038\u0072\u0075\u005a\u004f\u0076\u0072\u0053\u0063\u0073\u0066\u0032\u0051\u0065\u004e9\u002f\u0071r\u0055\u0047\u0038\u0071\u0045\u0062\u0055\u0057\u0064\u006f\u0045\u0059\u0071+\u000a\u006f\u0074\u0046\u004e\u0041\u0046N\u0078\u006c\u0047\u0062\u0078\u0062\u0044\u0048\u0063\u0064\u0047\u0056\u0061\u004d\u0030\u004f\u0058\u0064\u0058g\u0044y\u004c5\u0061\u0049\u0045\u0061\u0067\u004c\u0030\u0063\u0035\u0070\u0077\u006a\u0049\u0064\u0050G\u0049\u006e\u0034\u0036\u0066\u0037\u0038\u0065\u004d\u004a\u002b\u004a\u006b\u0064\u0063\u0070\u0044\n\u0044\u004a\u0061\u0071\u0059\u0058d\u0072\u007a5\u004b\u0065\u0073\u0068\u006aS\u0069\u0049\u0061\u0061\u0037\u006d\u0065\u006e\u0042\u0049\u0041\u0058\u0053\u0034\u0055\u0046\u0078N\u0066H\u0068\u004e\u0030\u0048\u0043\u0059\u005a\u0059\u0071\u0051\u0047\u0037\u0062K+\u0073\u0035\u0072R\u0048\u006f\u006e\u0079\u0064\u004eW\u0045\u0047\u000a\u0048\u0038M\u0079\u0076\u00722\u0070\u0079\u0061\u0032K\u0072\u004d\u0075m\u0066\u006d\u0041\u0078\u0055\u0042\u0036\u0066\u0065\u006e\u0043\u002f4\u004f\u0030\u0057\u00728\u0067\u0066\u0050\u004f\u0055\u0038R\u0069\u0074\u006d\u0062\u0044\u0076\u0051\u0050\u0049\u0052\u0058\u004fL\u0034\u0076\u0054B\u0072\u0042\u0064\u0062a\u0041\u000a9\u006e\u0077\u004e\u0050\u002b\u0069\u002f\u002f\u0032\u0030\u004d\u00542\u0062\u0078\u006d\u0065\u0057\u0042\u002b\u0067\u0070\u0063\u0045\u0068G\u0070\u0058\u005a7\u0033\u0033\u0061\u007a\u0051\u0078\u0072\u0043\u0033\u004a\u0034\u0076\u0033C\u005a\u006d\u0045\u004eS\u0074\u0044\u004b\u002f\u004b\u0044\u0053\u0050\u004b\u0055\u0047\u0066\u00756\u000a\u0066\u0077I\u0044\u0041\u0051\u0041\u0042\u000a\u002d\u002d\u002d\u002d\u002dE\u004e\u0044\u0020\u0050\u0055\u0042\u004c\u0049\u0043 \u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\n";const (LicenseTierUnlicensed ="\u0075\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";LicenseTierCommunity ="\u0063o\u006d\u006d\u0075\u006e\u0069\u0074y";LicenseTierIndividual ="\u0069\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c";LicenseTierBusiness ="\u0062\u0075\u0073\u0069\u006e\u0065\u0073\u0073";);

// LicenseKey represents a license key for UniOffice.
type LicenseKey struct{LicenseId string `json:"license_id"`;CustomerId string `json:"customer_id"`;CustomerName string `json:"customer_name"`;Tier string `json:"tier"`;CreatedAt _a .Time `json:"-"`;CreatedAtInt int64 `json:"created_at"`;ExpiresAt _a .Time `json:"-"`;ExpiresAtInt int64 `json:"expires_at"`;CreatorName string `json:"creator_name"`;CreatorEmail string `json:"creator_email"`;UniPDF bool `json:"unipdf"`;UniOffice bool `json:"unioffice"`;Trial bool `json:"trial"`;};