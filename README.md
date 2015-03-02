#i18n
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/rocwong/i18n)
[![GoCover](http://gocover.io/_badge/github.com/rocwong/i18n)](http://gocover.io/github.com/rocwong/i18n)

Smart internationalization for golang.

## Usage
~~~go
package main

import (
  "github.com/rocwong/i18n"
)

func main() {
  i18n.LoadMessages("testdata")
  // Return: Hello, i18n
  i18n.Translate("en", "greeting", "i18n")
  // Return: Hey, i18n
  i18n.Translate("en-GB", "greeting", "i18n")
  // Return: Howdy, i18n
  i18n.Translate("en-US", "greeting", "i18n")
}

~~~

## Message file
When creating new message files, there are a couple of rules to keep in mind:

- The file extension determines the language of the message file and should be an [ISO 639-1 code](http://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
- Each message file is effectively a [robfig/config](https://github.com/robfig/config) and supports all [robfig/config](https://github.com/robfig/config) features.
- Message files should be UTF-8 encoded. While this is not a hard requirement, it is best practice.
- Must be one blank line at the end of the message file.

#### Organizing message files
There are no restrictions on message file names, a message file name can be anything as long as it has a valid extention. you can free to organize the message files however you want.

For example, you may want to take a traditional approach and define 1 single message file per language:
~~~go
/app
    /i18n
        i18n.en
        i18n.zh
        ...
~~~

Another approach would be to create multiple files for the same language and organize them based on the kind of messages they contain:
~~~go
/app
    /i18n
        user.en
        admin.en
        user.zh
        admin.zh
        ...
~~~

## Regions
Region-specific messages should be defined in sections with the same name. For example, suppose that we want to greet all English speaking users with "`Hello`", all British users with "`Hey`" and all American users with "`Howdy`". In order to accomplish this, we could define the following message file `greeting.en`:

~~~go
greeting=Hello

[GB]
greeting=Hey

[US]
greeting=Howdy
// Must be one blank line at the end of the message file.
~~~
For users who have defined English (`en`) as their preferred language, i18n would resolve `greeting` to `Hello`. Only in specific cases where the user’s locale has been explicitly defined as `en-GB` or `en-US` would the greeting message be resolved using the specific sections.


