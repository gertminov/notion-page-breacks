# notion-page-breacks
this is a little script that allows you to manually decide where a page should break when exporting a notion document.  \
It adds `page-break-after: always` and `visibility: hidden !important;` to `<hr>` tags

## Usage
1. Add a seperator `---` in notion where you want a new page to begin
2. Export the notion page as HTML
3. Extract the .zip File
4. Run `notion-page-breacks filename.html`
5. Open the new HTML file in the Browser and hit `CTRL + P` to open the print dialog
6. Choose "Save as PDF" and hit print

You should now have a beatiful PDF file that has page breacks, where you want them
