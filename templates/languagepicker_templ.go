// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// class="block appearance-none w-full border-0 border-gray-200 text-gray-700 py-1.5 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
func Langpicker() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form class=\"w-full max-w-sm\" hx-post=\"bibleselect\" hx-target=\"#biblepicker\"><input id=\"defaultlang\" type=\"hidden\" value=\"en\" name=\"defaultlang\"><div class=\"inline-block relative w-25\"><select id=\"langselect\" name=\"langselect\" class=\"block appearance-none w-full bg-white border-0 border-gray-400 hover:border-gray-500 px-4 py-2 pr-8 rounded leading-tight focus:outline-none focus:shadow-outline\" onchange=\"submitandclearlang()\"><option disabled selected value=\"Language\">Language</option> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range Languages() {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(item.id))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(item.name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/languagepicker.templ`, Line: 15, Col: 42}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select><div class=\"pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700\"><svg class=\"fill-current h-4 w-4\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\"><path d=\"M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z\"></path></svg></div></div><button id=\"langpickerclick\" type=\"submit\" class=\"invisible\" type=\"button\" hidden>Submit</button></form><script>\n\tlangconversion = { \"aa\": \"aar\", \"ab\": \"abk\", \"af\": \"afr\", \"ak\": \"aka\", \"sq\": \"alb\", \"am\": \"amh\", \"ar\": \"ara\", \"an\": \"arg\", \"hy\": \"hye\", \"as\": \"asm\", \"av\": \"ava\", \"ae\": \"ave\", \"ay\": \"aym\", \"az\": \"aze\", \"ba\": \"bak\", \"bm\": \"bam\", \"eu\": \"eus\", \"be\": \"bel\", \"bn\": \"ben\", \"bh\": \"bih\", \"bi\": \"bis\", \"bo\": \"tib\", \"bs\": \"bos\", \"br\": \"bre\", \"bg\": \"bul\", \"my\": \"mya\", \"ca\": \"cat\", \"cs\": \"cze\", \"ch\": \"cha\", \"ce\": \"che\", \"zh\": \"zho\", \"cu\": \"chu\", \"cv\": \"chv\", \"kw\": \"cor\", \"co\": \"cos\", \"cr\": \"cre\", \"cy\": \"wel\", \"da\": \"dan\", \"de\": \"ger\", \"dv\": \"div\", \"nl\": \"dut\", \"dz\": \"dzo\", \"el\": \"gre\", \"en\": \"eng\", \"eo\": \"epo\", \"et\": \"est\", \"ee\": \"ewe\", \"fo\": \"fao\", \"fa\": \"per\", \"fj\": \"fij\", \"fi\": \"fin\", \"fr\": \"fra\", \"fy\": \"fry\", \"ff\": \"ful\", \"ka\": \"geo\", \"gd\": \"gla\", \"ga\": \"gle\", \"gl\": \"glg\", \"gv\": \"glv\", \"gn\": \"grn\", \"gu\": \"guj\", \"ht\": \"hat\", \"ha\": \"hau\", \"he\": \"heb\", \"hz\": \"her\", \"hi\": \"hin\", \"ho\": \"hmo\", \"hr\": \"hrv\", \"hu\": \"hun\", \"ig\": \"ibo\", \"is\": \"ice\", \"io\": \"ido\", \"ii\": \"iii\", \"iu\": \"iku\", \"ie\": \"ile\", \"ia\": \"ina\", \"id\": \"ind\", \"ik\": \"ipk\", \"it\": \"ita\", \"jv\": \"jav\", \"ja\": \"jpn\", \"kl\": \"kal\", \"kn\": \"kan\", \"ks\": \"kas\", \"kr\": \"kau\", \"kk\": \"kaz\", \"km\": \"khm\", \"ki\": \"kik\", \"rw\": \"kin\", \"ky\": \"kir\", \"kv\": \"kom\", \"kg\": \"kon\", \"ko\": \"kor\", \"kj\": \"kua\", \"ku\": \"kur\", \"lo\": \"lao\", \"la\": \"lat\", \"lv\": \"lav\", \"li\": \"lim\", \"ln\": \"lin\", \"lt\": \"lit\", \"lb\": \"ltz\", \"lu\": \"lub\", \"lg\": \"lug\", \"mk\": \"mkd\", \"mh\": \"mah\", \"ml\": \"mal\", \"mi\": \"mri\", \"mr\": \"mar\", \"ms\": \"may\", \"mg\": \"mlg\", \"mt\": \"mlt\", \"mn\": \"mon\", \"na\": \"nau\", \"nv\": \"nav\", \"nr\": \"nbl\", \"nd\": \"nde\", \"ng\": \"ndo\", \"ne\": \"nep\", \"nn\": \"nno\", \"nb\": \"nob\", \"no\": \"nor\", \"ny\": \"nya\", \"oc\": \"oci\", \"oj\": \"oji\", \"or\": \"ori\", \"om\": \"orm\", \"os\": \"oss\", \"pa\": \"pan\", \"pi\": \"pli\", \"pl\": \"pol\", \"pt\": \"por\", \"ps\": \"pus\", \"qu\": \"que\", \"rm\": \"roh\", \"ro\": \"ron\", \"rn\": \"run\", \"ru\": \"rus\", \"sg\": \"sag\", \"sa\": \"san\", \"si\": \"sin\", \"sk\": \"slk\", \"sl\": \"slv\", \"se\": \"sme\", \"sm\": \"smo\", \"sn\": \"sna\", \"sd\": \"snd\", \"so\": \"som\", \"st\": \"sot\", \"es\": \"spa\", \"sc\": \"srd\", \"sr\": \"srp\", \"ss\": \"ssw\", \"su\": \"sun\", \"sw\": \"swa\", \"sv\": \"swe\", \"ty\": \"tah\", \"ta\": \"tam\", \"tt\": \"tat\", \"te\": \"tel\", \"tg\": \"tgk\", \"tl\": \"tgl\", \"th\": \"tha\", \"ti\": \"tir\", \"to\": \"ton\", \"tn\": \"tsn\", \"ts\": \"tso\", \"tk\": \"tuk\", \"tr\": \"tur\", \"tw\": \"twi\", \"ug\": \"uig\", \"uk\": \"ukr\", \"ur\": \"urd\", \"uz\": \"uzb\", \"ve\": \"ven\", \"vi\": \"vie\", \"vo\": \"vol\", \"wa\": \"wln\", \"wo\": \"wol\", \"xh\": \"xho\", \"yi\": \"yid\", \"yo\": \"yor\", \"za\": \"zha\", \"zu\": \"zul\" };\n\tfunction submitandclearlang() {\n\t\tdocument.getElementById('langpickerclick').click(); \n\t\tdocument.getElementById('biblecontent').innerHTML = '';\n\t\tdocument.getElementById('bookpicker').innerHTML = '';\n\t\tdocument.getElementById('chapterpicker').innerHTML = '';\n\t}\t\n\tfunction setValue() {\n\t\tif (navigator.language.includes('-')) document.getElementById('langselect').value = langconversion[navigator.language.split(\"-\")[0]]\n\t\telse document.getElementById('langselect').value = langconversion[navigator.language]\n\t\tif (document.getElementById('langselect').value == 'Language' || document.getElementById('langselect').value == '') {\n\t\t\tconsole.log(\"Language Couldn't be determined, defaulting to English\")\n\t\t\tdocument.getElementById('langselect').value = 'eng'\n\t\t}\n\t}\n\n\twindow.onload = async function(){\n\t\tawait setValue()\n\t\tdocument.getElementById('langpickerclick').click()\n\t}\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
