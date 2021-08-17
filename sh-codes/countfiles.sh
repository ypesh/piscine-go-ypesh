#ls | grep -v ^l |wc -l
#./countfiles.sh | cat -e
find -type f -o -type d | wc -l
