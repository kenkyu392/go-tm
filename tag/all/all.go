package english

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
ab-RU,Abkhaz (Russia)
ae-IR,Avestan (Islamic Republic of Iran)
ak-GH,Akan (Republic of Ghana)
am-ET,Amharic (Ethiopia)
as-IN,Assamese (India)
ast-ES,Asturian (Spain)
ay-BO,Aymara (Bolivia)
az-AZ,Azerbaijani (Azerbaijan)
ba-RU,Bashkir (Russia)
be-BY,Belarusian (Belarus)
bg-BG,Bulgarian (Bulgaria)
bh-IN,Bihari (India)
bi-VU,Bislama (Vanuatu)
bn-IN,Bengali (India)
br-FR,Breton (France)
bs-BA,Bosnian (Bosnia and Herzegovina)
ce-RU,Chechen (Russia)
co-FR,Corsican (France)
cs-CZ,Czech (Czech Republic)
cv-RU,Chuvash (Russia)
cy-GB,Welsh (United Kingdom)
da-DK,Danish (Denmark)
dsb-DE,Lower Sorbian (Germany)
dv-IN,Divehi (India)
dv-MV,Dihevi (Maldives)
dz-BT,Dzongkha (Bhutan)
el-CY,Greek (Cyprus)
el-GR,Greek (Greece)
et-EE,Estonian (Estonia)
eu-ES,Basque (Spain)
fa-IR,Farsi/Persian (Islamic Republic of Iran)
fi-FI,Finnish (Finland)
fj-FJ,Fijian (Fiji)
fy-NL,Frisian (Netherlands)
ga-IE,Irish (Ireland)
gd-GB,Scottish Gaelic (United Kingdom)
gl-ES,Galician (Spain)
gn-PY,Guarani (Paraguay)
grc-GR,Ancient Greek (Greece)
gu-IN,Gujarati (India)
gv-IM,Manx (Isle of Man)
he-IL,Hebrew (Israel)
hi-IN,Hindi (India)
ho-PG,Hiri Motu (Independent State of Papua New Guinea)
hr-BA,Croatian (Bosnia and Herzegovina)
hr-HR,Croatian (Croatia)
hsb-DE,Upper Sorbian (Germany)
ht-HT,Haitian (Republic of Haiti)
hu-HU,Hungarian (Hungary)
hy-AM,Armenian (Armenia)
hz-NA,Herero (Republic of Namibia)
id-ID,Indonesian (Indonesia)
ik-US,Inupiaq (United States of America)
is-IS,Icelandic (Iceland)
iu-CA,Inuktitut (Canada)
jv-ID,Javanese (Indonesia)
ka-GE,Georgian (Georgia)
kg-CG,Kongo (Republic of the Congo)
kk-KZ,Kazakh (Kazakhstan)
kl-GL,Kalaallisut/Greenlandic (Greenland)
km-KH,Khmer (Cambodia)
kn-IN,Kannada (India)
kok-IN,Konkani (India)
ks-IN,Kashmiri (India)
ku-IQ,Kurdish (Iraq)
ku-IR,Kurdish (Islamic Republic of Iran)
kv-RU,Komi (Russia)
kw-GB,Cornish (United Kingdom)
ky-KG,Kyrgyz (Kyrgyz Republic)
la-VA,Latin (Vatican City State)
lb-LU,Luxembourgish (Luxembourg)
li-LU,Limburgish (Luxembourg)
lo-LA,Lao (Lao People's Democratic Republic)
lt-LT,Lithuanian (Lithuania)
lv-LV,Latvian (Latvia)
mg-MG,Malagasy (Madagascar)
mh-MH,Marshallese (Republic of the Marshall Islands)
mi-NZ,Maori (New Zealand)
mk-MK,Macedonian (North Macedonia)
ml-IN,Malayalam (India)
mn-MN,Mongolian (Mongolia)
mo-MD,Moldovan (Republic of Moldova)
mr-IN,Marathi (India)
mt-MT,Maltese (Malta)
my-MM,Burmese (Myanmar)
na-NR,Nauruan (Republic of Nauru)
nb-NO,Norwegian (Bokmål)
nd-ZW,North Ndebele (Zimbabwe)
ne-IN,Nepali (India)
ne-NP,Nepali (Nepal)
ng-NA,Ndonga (Republic of Namibia)
nn-NO,Norwegian (Nynorsk)
nso-ZA,Northern Sotho (South Africa)
ny-MW,Chewa (Malawi)
oc-FR,Occitan (France)
om-ET,Oromo (Ethiopia)
om-KE,Oromo (Kenya)
or-IN,Oriya (India)
os-GE,Ossetian (Georgia)
os-RU,Ossetian (Russia)
pa-IN,Panjabi/Punjabi (India)
pi-IN,Pali (India)
pl-PL,Polish (Poland)
ps-AF,Pushto/Pashto (Afghanistan)
rm-CH,Romansh (Switzerland)
rn-BI,Rundi (Burundi)
rw-RW,Kinyarwanda (Rwanda)
sa-IN,Sanskrit (India)
sc-IT,Sardinian (Italy)
sd-PK,Sindhi (Islamic Republic of Pakistan)
sg-CF,Sango (Central African Republic)
si-IN,Sinhalese (India)
si-LK,Sinhala (Democratic Socialist Republic of Sri Lanka)
sk-SK,Slovak (Slovakia)
sl-SI,Slovenian (Slovenia)
sm-WS,Samoan (Independent State of Samoa)
sn-ZW,Shona (Zimbabwe)
su-ID,Sundanese (Indonesia)
syr-TR,Syriac (Syria)
ta-IN,Tamil (India)
te-IN,Telugu (India)
tet-ID,Tetum (Indonesia)
tet-TL,Tetum (Democratic Republic of Timor-Leste)
tg-TJ,Tajik (Tajikistan)
th-TH,Thai (Thailand)
tk-TM,Turkmen (Turkmenistan)
tl-PH,Tagalog (Philippines)
to-TO,Tongan (Tonga)
ts-ZA,Tsonga (South Africa)
tt-RU,Tatar (Russia)
tw-GH,Twi (Ghana)
ty-PF,Tahitian (French Polynesia)
ug-CN,Uyghur (People's Republic of China)
uk-UA,Ukrainian (Ukraine)
ur-IN,Urdu (India)
ur-PK,Urdu (Islamic Republic of Pakistan)
uz-AF,Uzbek (Afghanistan)
uz-UZ,Uzbek (Uzbekistan)
vi-VN,Vietnamese (Vietnam)
wa-BE,Walloon (Belgium)
wo-SN,Wolof (Senegal)
xh-ZA,Xhosa (South Africa)
yi-BA,Yiddish (Bosnia and Herzegovina)
zu-ZA,Zulu (South Africa)
*/

func init() {
	tag.Register("ab-RU", "Abkhaz (Russia)")
	tag.Register("ae-IR", "Avestan (Islamic Republic of Iran)")
	tag.Register("ak-GH", "Akan (Republic of Ghana)")
	tag.Register("am-ET", "Amharic (Ethiopia)")
	tag.Register("as-IN", "Assamese (India)")
	tag.Register("ast-ES", "Asturian (Spain)")
	tag.Register("ay-BO", "Aymara (Bolivia)")
	tag.Register("az-AZ", "Azerbaijani (Azerbaijan)")
	tag.Register("ba-RU", "Bashkir (Russia)")
	tag.Register("be-BY", "Belarusian (Belarus)")
	tag.Register("bg-BG", "Bulgarian (Bulgaria)")
	tag.Register("bh-IN", "Bihari (India)")
	tag.Register("bi-VU", "Bislama (Vanuatu)")
	tag.Register("bn-IN", "Bengali (India)")
	tag.Register("br-FR", "Breton (France)")
	tag.Register("bs-BA", "Bosnian (Bosnia and Herzegovina)")
	tag.Register("ce-RU", "Chechen (Russia)")
	tag.Register("co-FR", "Corsican (France)")
	tag.Register("cs-CZ", "Czech (Czech Republic)")
	tag.Register("cv-RU", "Chuvash (Russia)")
	tag.Register("cy-GB", "Welsh (United Kingdom)")
	tag.Register("da-DK", "Danish (Denmark)")
	tag.Register("dsb-DE", "Lower Sorbian (Germany)")
	tag.Register("dv-IN", "Divehi (India)")
	tag.Register("dv-MV", "Dihevi (Maldives)")
	tag.Register("dz-BT", "Dzongkha (Bhutan)")
	tag.Register("el-CY", "Greek (Cyprus)")
	tag.Register("el-GR", "Greek (Greece)")
	tag.Register("et-EE", "Estonian (Estonia)")
	tag.Register("eu-ES", "Basque (Spain)")
	tag.Register("fa-IR", "Farsi/Persian (Islamic Republic of Iran)")
	tag.Register("fi-FI", "Finnish (Finland)")
	tag.Register("fj-FJ", "Fijian (Fiji)")
	tag.Register("fy-NL", "Frisian (Netherlands)")
	tag.Register("ga-IE", "Irish (Ireland)")
	tag.Register("gd-GB", "Scottish Gaelic (United Kingdom)")
	tag.Register("gl-ES", "Galician (Spain)")
	tag.Register("gn-PY", "Guarani (Paraguay)")
	tag.Register("grc-GR", "Ancient Greek (Greece)")
	tag.Register("gu-IN", "Gujarati (India)")
	tag.Register("gv-IM", "Manx (Isle of Man)")
	tag.Register("he-IL", "Hebrew (Israel)")
	tag.Register("hi-IN", "Hindi (India)")
	tag.Register("ho-PG", "Hiri Motu (Independent State of Papua New Guinea)")
	tag.Register("hr-BA", "Croatian (Bosnia and Herzegovina)")
	tag.Register("hr-HR", "Croatian (Croatia)")
	tag.Register("hsb-DE", "Upper Sorbian (Germany)")
	tag.Register("ht-HT", "Haitian (Republic of Haiti)")
	tag.Register("hu-HU", "Hungarian (Hungary)")
	tag.Register("hy-AM", "Armenian (Armenia)")
	tag.Register("hz-NA", "Herero (Republic of Namibia)")
	tag.Register("id-ID", "Indonesian (Indonesia)")
	tag.Register("ik-US", "Inupiaq (United States of America)")
	tag.Register("is-IS", "Icelandic (Iceland)")
	tag.Register("iu-CA", "Inuktitut (Canada)")
	tag.Register("jv-ID", "Javanese (Indonesia)")
	tag.Register("ka-GE", "Georgian (Georgia)")
	tag.Register("kg-CG", "Kongo (Republic of the Congo)")
	tag.Register("kk-KZ", "Kazakh (Kazakhstan)")
	tag.Register("kl-GL", "Kalaallisut/Greenlandic (Greenland)")
	tag.Register("km-KH", "Khmer (Cambodia)")
	tag.Register("kn-IN", "Kannada (India)")
	tag.Register("kok-IN", "Konkani (India)")
	tag.Register("ks-IN", "Kashmiri (India)")
	tag.Register("ku-IQ", "Kurdish (Iraq)")
	tag.Register("ku-IR", "Kurdish (Islamic Republic of Iran)")
	tag.Register("kv-RU", "Komi (Russia)")
	tag.Register("kw-GB", "Cornish (United Kingdom)")
	tag.Register("ky-KG", "Kyrgyz (Kyrgyz Republic)")
	tag.Register("la-VA", "Latin (Vatican City State)")
	tag.Register("lb-LU", "Luxembourgish (Luxembourg)")
	tag.Register("li-LU", "Limburgish (Luxembourg)")
	tag.Register("lo-LA", "Lao (Lao People's Democratic Republic)")
	tag.Register("lt-LT", "Lithuanian (Lithuania)")
	tag.Register("lv-LV", "Latvian (Latvia)")
	tag.Register("mg-MG", "Malagasy (Madagascar)")
	tag.Register("mh-MH", "Marshallese (Republic of the Marshall Islands)")
	tag.Register("mi-NZ", "Maori (New Zealand)")
	tag.Register("mk-MK", "Macedonian (North Macedonia)")
	tag.Register("ml-IN", "Malayalam (India)")
	tag.Register("mn-MN", "Mongolian (Mongolia)")
	tag.Register("mo-MD", "Moldovan (Republic of Moldova)")
	tag.Register("mr-IN", "Marathi (India)")
	tag.Register("mt-MT", "Maltese (Malta)")
	tag.Register("my-MM", "Burmese (Myanmar)")
	tag.Register("na-NR", "Nauruan (Republic of Nauru)")
	tag.Register("nb-NO", "Norwegian (Bokmål)")
	tag.Register("nd-ZW", "North Ndebele (Zimbabwe)")
	tag.Register("ne-IN", "Nepali (India)")
	tag.Register("ne-NP", "Nepali (Federal Democratic Republic of Nepal)")
	tag.Register("ng-NA", "Ndonga (Republic of Namibia)")
	tag.Register("nn-NO", "Norwegian (Nynorsk)")
	tag.Register("nso-ZA", "Northern Sotho (South Africa)")
	tag.Register("ny-MW", "Chewa (Malawi)")
	tag.Register("oc-FR", "Occitan (France)")
	tag.Register("om-ET", "Oromo (Ethiopia)")
	tag.Register("om-KE", "Oromo (Kenya)")
	tag.Register("or-IN", "Oriya (India)")
	tag.Register("os-GE", "Ossetian (Georgia)")
	tag.Register("os-RU", "Ossetian (Russia)")
	tag.Register("pa-IN", "Panjabi/Punjabi (India)")
	tag.Register("pi-IN", "Pali (India)")
	tag.Register("pl-PL", "Polish (Poland)")
	tag.Register("ps-AF", "Pushto/Pashto (Afghanistan)")
	tag.Register("rm-CH", "Romansh (Switzerland)")
	tag.Register("rn-BI", "Rundi (Burundi)")
	tag.Register("rw-RW", "Kinyarwanda (Rwanda)")
	tag.Register("sa-IN", "Sanskrit (India)")
	tag.Register("sc-IT", "Sardinian (Italy)")
	tag.Register("sd-PK", "Sindhi (Islamic Republic of Pakistan)")
	tag.Register("sg-CF", "Sango (Central African Republic)")
	tag.Register("si-IN", "Sinhalese (India)")
	tag.Register("si-LK", "Sinhala (Democratic Socialist Republic of Sri Lanka)")
	tag.Register("sk-SK", "Slovak (Slovakia)")
	tag.Register("sl-SI", "Slovenian (Slovenia)")
	tag.Register("sm-WS", "Samoan (Independent State of Samoa)")
	tag.Register("sn-ZW", "Shona (Zimbabwe)")
	tag.Register("su-ID", "Sundanese (Indonesia)")
	tag.Register("syr-TR", "Syriac (Syria)")
	tag.Register("ta-IN", "Tamil (India)")
	tag.Register("te-IN", "Telugu (India)")
	tag.Register("tet-ID", "Tetum (Indonesia)")
	tag.Register("tet-TL", "Tetum (Democratic Republic of Timor-Leste)")
	tag.Register("tg-TJ", "Tajik (Tajikistan)")
	tag.Register("th-TH", "Thai (Thailand)")
	tag.Register("tk-TM", "Turkmen (Turkmenistan)")
	tag.Register("tl-PH", "Tagalog (Philippines)")
	tag.Register("to-TO", "Tongan (Tonga)")
	tag.Register("ts-ZA", "Tsonga (South Africa)")
	tag.Register("tt-RU", "Tatar (Russia)")
	tag.Register("tw-GH", "Twi (Ghana)")
	tag.Register("ty-PF", "Tahitian (French Polynesia)")
	tag.Register("ug-CN", "Uyghur (People's Republic of China)")
	tag.Register("uk-UA", "Ukrainian (Ukraine)")
	tag.Register("ur-IN", "Urdu (India)")
	tag.Register("ur-PK", "Urdu (Islamic Republic of Pakistan)")
	tag.Register("uz-AF", "Uzbek (Afghanistan)")
	tag.Register("uz-UZ", "Uzbek (Uzbekistan)")
	tag.Register("vi-VN", "Vietnamese (Vietnam)")
	tag.Register("wa-BE", "Walloon (Belgium)")
	tag.Register("wo-SN", "Wolof (Senegal)")
	tag.Register("xh-ZA", "Xhosa (South Africa)")
	tag.Register("yi-BA", "Yiddish (Bosnia and Herzegovina)")
	tag.Register("zu-ZA", "Zulu (South Africa)")
}
