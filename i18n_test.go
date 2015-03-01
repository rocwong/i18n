package i18n

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Translate(t *testing.T) {
	LoadMessages("testdata")
	Convey("Translates content to target language", t, func() {
		So(Translate("en", "greeting", "i18n"), ShouldEqual, "Hello, i18n")
		So(Translate("en-GB", "greeting", "i18n"), ShouldEqual, "Hey, i18n")
		So(Translate("en-US", "greeting", "i18n"), ShouldEqual, "Howdy, i18n")
		So(Translate("en", "file"), ShouldEqual, "test/b.en")

		So(Translate("zh", "greeting", "i18n"), ShouldEqual, "你好, i18n")

		So(Translate("zh", "job", "i18n"), ShouldEqual, "job")
		So(Translate("ja", "greeting"), ShouldEqual, "greeting")

	})
	Convey("Gets all currently loaded message languages", t, func() {
		langs := ListLanguages()
		So(len(langs), ShouldEqual, 2)
	})
}
