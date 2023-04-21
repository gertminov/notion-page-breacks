# notion-page-breacks
this is a little tool that allows you to manually decide where a page should break when exporting a notion document.  \
It adds `page-break-after: always` and `visibility: hidden !important;` to `<hr>` tags

## Install
Download the newest version from [the releases page](https://github.com/gertminov/notion-page-breacks/releases), 
and place the `.exe`in a directory that is in the PATH or copy it in the directory where you exportet the notion pages

## Usage
1. Add a separator `---` in notion where you want a new page to begin
2. Export the notion page as HTML
3. Open the Terminal and run `notion-pb "./filename.zip"` \
(optional you can use `glob`patterns to add multiple files like `"./*.html"` to convert all `.html` files in the current folder)
4. Open the new HTML file in `./fixed/` in your Browser and hit `CTRL + P` to open the print dialog
5. Choose "Save as PDF" and hit print

You should now have a beautiful PDF file that has page breaks, where you want them


## Options
```
-o [outputname]            Give a custom output name.
<input>.html                 adds line breacks to extraced .html file
<input>.zip                  extracts archive and adds line breacks
```
