# nvdslurp

download a cpe.gz from : https://nvd.nist.gov/products/cpe

e.g: download [this](https://nvd.nist.gov/feeds/xml/cpe/dictionary/official-cpe-dictionary_v2.3.xml.gz) file.

For now pipe it to [fzf](https://github.com/junegunn/fzf)

Run `./nvdslurp -in=/path/to/official-cpe-dictionary_v2.3.xml.gz | fzf`
