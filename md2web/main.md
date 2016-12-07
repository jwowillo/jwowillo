## md2web

md2web is a markdown website generator.

### Installation

Run `git clone https://github.com/jwowillo/md2web.git`.

Run `make`.

### Running

Change to a folder containing other folders and markdown files. Run `md2web
<?port>`. A website based on the folder's contents will now be running at the
optionally specified port.

### Use

By default, md2web serves markdown files based on the folder structure they
exist in. Static files can be included in the markdown files by linking to files
in a folder called `static` with the link
`http://cdn.<domain>.<tld>/<file>`. The favicon served can be set by placing a
file at `static/favicon.png`. To customize HTML injected into different sections
of the site, files can be placed into folders at `static/head`, `static/header`,
`static/nav`, `static/aside`, and `static/section`. These will be prepended to
the corresponding parts of the template.

### Mancalai Integration Plan

Steps:
* Configure static subdomain, api subdomain, and static path prefix in mancalai.
* Git clone mancalai into `apps` directory in Makefile.
* Recursively call all Makefiles in `apps`.
* Automatically register subdomains using dnsimple.
