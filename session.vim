let SessionLoad = 1
let s:so_save = &g:so | let s:siso_save = &g:siso | setg so=0 siso=0 | setl so=-1 siso=-1
let v:this_session=expand("<sfile>:p")
silent only
silent tabonly
cd ~/Documents/GitHub/flareup
if expand('%') == '' && !&modified && line('$') <= 1 && getline(1) == ''
  let s:wipebuf = bufnr('%')
endif
let s:shortmess_save = &shortmess
if &shortmess =~ 'A'
  set shortmess=aoOA
else
  set shortmess=aoO
endif
badd +1 ~/Documents/GitHub/flareup
badd +1 main.go
badd +1 internal/cloudflare/structs.go
badd +119 internal/cloudflare/cloudflare.go
argglobal
%argdel
$argadd ~/Documents/GitHub/flareup
edit internal/cloudflare/cloudflare.go
argglobal
setlocal foldmethod=manual
setlocal foldexpr=0
setlocal foldmarker={{{,}}}
setlocal foldignore=#
setlocal foldlevel=99
setlocal foldminlines=1
setlocal foldnestmax=20
setlocal foldenable
silent! normal! zE
3,12fold
19,20fold
25,26fold
17,30fold
33,34fold
44,47fold
42,48fold
52,54fold
55,56fold
51,57fold
37,59fold
66,68fold
65,69fold
62,70fold
80,81fold
85,86fold
73,89fold
98,99fold
103,104fold
109,110fold
114,115fold
112,116fold
92,118fold
128,129fold
134,135fold
140,141fold
143,144fold
121,146fold
154,161fold
167,168fold
149,173fold
let &fdl = &fdl
let s:l = 119 - ((40 * winheight(0) + 31) / 62)
if s:l < 1 | let s:l = 1 | endif
keepjumps exe s:l
normal! zt
keepjumps 119
normal! 0
lcd ~/Documents/GitHub/flareup
tabnext 1
if exists('s:wipebuf') && len(win_findbuf(s:wipebuf)) == 0 && getbufvar(s:wipebuf, '&buftype') isnot# 'terminal'
  silent exe 'bwipe ' . s:wipebuf
endif
unlet! s:wipebuf
set winheight=1 winwidth=20
let &shortmess = s:shortmess_save
let s:sx = expand("<sfile>:p:r")."x.vim"
if filereadable(s:sx)
  exe "source " . fnameescape(s:sx)
endif
let &g:so = s:so_save | let &g:siso = s:siso_save
doautoall SessionLoadPost
unlet SessionLoad
" vim: set ft=vim :
