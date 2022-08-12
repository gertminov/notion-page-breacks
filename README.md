# notion-page-breacks
this is a little script that allows you to manually decide where a page should break when exporting a notion document.  \
It adds `page-break-after: always` and `visibility: hidden !important;` to `<hr>` tags

## Usage

### With UI

1. Add a separator `---` in notion where you want a new page to begin
2. Export the notion page as HTML
3. start `notion-page-breaks`
4. Select the `.zip` file the notion generated in the file picker
5. Click `yes` if the program asks you, wheter you want to open the new file in the browser
6. Open the new HTML file in `./extract` the Browser and hit `CTRL + P` to open the print dialog
7. Choose "Save as PDF" and hit print

### Without UI
1. Add a separator `---` in notion where you want a new page to begin
2. Export the notion page as HTML
3. Open the Terminal and run `notion-page-breaks filename.zip`
4. Open the new HTML file in `./extract` the Browser and hit `CTRL + P` to open the print dialog
5. Choose "Save as PDF" and hit print

You should now have a beautiful PDF file that has page breaks, where you want them


## Options
```
-o [outputname]            Give a custom output name.
<input>.html                 adds line breacks to extraced .html file
<input>.zip                  extracts archive and adds line breacks
```