find . -type f \( -name "*.sh" \) | sed 's|^./||'| sed 's/.sh//g' | rev | cut -d '/' -f 1 | rev | sort -r
