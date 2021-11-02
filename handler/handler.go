package handler

import (
"html/template"
"log"
"net/http"
"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
log.Printf(r.URL.Path)

if r.URL.Path != "/" {
http.NotFound(w, r)
return
}

tmpl, err := template.ParseFiles(path.Join("views", "index.html"),path.Join("views", "header.html"),path.Join("views", "footer.html"),path.Join("views", "navbar.html"),path.Join("views", "section1.html"),path.Join("views", "section2.html"),path.Join("views", "services.html"),path.Join("views", "cinema1.html"),path.Join("views", "cinema2.html"),path.Join("views", "cinema3.html"),path.Join("views", "contact.html"),path.Join("views", "banner.html"))
if err != nil {
log.Println(err)
http.Error(w, "oopsss.... error pages", http.StatusInternalServerError)
return 
}

data := map[string]interface{}{
"logo": "https://1.bp.blogspot.com/-6B-s5iMYeNw/YQ16KQUwxcI/AAAAAAAAQMc/T2nmdoYWIP0vJiTEEr3EFGuuQxtx67EsACLcBGAsYHQ/s729/online.png",
"title": "Axcora", 
"description": "Layanan pembuatan website modern dan aplikasi android",
"menu1": "home",
"menu2": "about",
"menu3": "technology",
"menu4": "info",
"menu5": "contact",
"youtube": "SZ77qUZX7i4",
"youtube1": "oyBAPnwIlLE",
"youtube2": "DK756CZOes0",
"image1": "https://1.bp.blogspot.com/-NSo3h1NqtLw/YQ4GM17__9I/AAAAAAAAQN0/-IWsuvIae60luAgM3c4qprrYkITI_P44gCLcBGAsYHQ/s630/faster.png",
"section1": "Pastikan kebutuhan website modern mu dan aplikasi android mu bersama kami, axcora technology memberikan solusi terbaik dalam kebutuhan development webapp project, baik untuk company profile, shop store outlet , restoran cafe rumah makan hingga sekolah. Dengan berbagai technology dalam kebutuhan deployment menjadi lebih baik untuk memberikan kebebasan mu memilih technology favoritmu bersama kami, dengan memiliki website informasi maka kini lebih mudah untuk pelanggan cek informasi via web offical resmi mu. dengan single page application / progresive web app tech maka sekali produksi untuk kebutuhan web modern dan aplikasi android mu.",
"textinfo": "MULTI TECHNOLOGY",
"imginfo1": "https://mesinkasironline.web.app/img/createwebsiteusingangular.png",
"textinfo1": "Angular Web Apps",
"imginfo2": "https://mesinkasironline.web.app/img/createwebsiteusingreact.png",
"textinfo2": "React Web Apps",
"imginfo3": "https://svelte.dev/svelte-logo-outline.svg",
"textinfo3": "Svelte Web Apps",
"imginfo4": "https://mesinkasironline.web.app/icons/icon-384x384.png",
"textinfo4": "Gatsby Popular Web",
"imginfo5": "https://mesinkasironline.web.app/img/createwebsitestaticwithjekyll.png",
"textinfo5": "Jekyll Static Site",
"imginfo6": "https://mesinkasironline.web.app/img/createwebsitelaravel.png",
"textinfo6": "Laravel Web Apps",
"image2": "https://1.bp.blogspot.com/-Cur5BSTO61c/YQ2E8dSk9BI/AAAAAAAAQNo/GI4CsS7wT8c3tVpthlcM-9A_uLPlaL2PQCLcBGAsYHQ/s733/cf2.png",
"section2": "Dengan technology modern menjadi lebih baik dalam pembuatan website modern mu, ya kami menggunakan static site generasi terbaru dalam technology untuk project pembuatan website, selain itu dengan technology ini maka dapat digunakan juga untuk build aplikasi android mu, tentunya akan semakin mengirit anggaran belanja dalam build website modern dan aplikasi android all in one, selain itu jika dibutuhkan axcora tech juga menyediakan web app, website dan aplikasi pembukuan terintergasi untuk menunjang berbagai kebutuhan mu, saatnya revolusi digital dimulai bersama kami.",
"cover": "https://1.bp.blogspot.com/-wQ1WGUVYn_Y/YQ1_KkhvUYI/AAAAAAAAQMo/AkCE-EjfM7sJrusT_CFeoaij59aZ1TxRwCLcBGAsYHQ/s731/feat9.png",
"company": "AXCORA TECHNOLOGY",
"alamat": "Ruko Pasar Wisata Juanda Q.07",
"negara": "Sidoarjo - East Java - Indonesia",
"phone": "Whatsapp : 0856461044747 / Call : 0895339403223",

}

err = tmpl.Execute(w, data)
if err != nil {
log.Println(err)
http.Error(w, "oopsss.... error pages", http.StatusInternalServerError)
return 
}

}
