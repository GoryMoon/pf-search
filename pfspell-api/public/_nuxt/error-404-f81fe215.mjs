import{a as n,b as r,o as c,e as d,f as e,t as a,h as i,w as l,i as p,j as u,p as x,k as _}from"./entry-b5bab3ce.mjs";const m=t=>(x("data-v-bed9aeba"),t=t(),_(),t),f={class:"font-sans antialiased bg-white dark:bg-black text-black dark:text-white grid min-h-screen place-content-center overflow-hidden"},h=m(()=>e("div",{class:"fixed left-0 right-0 spotlight z-10"},null,-1)),g={class:"max-w-520px text-center z-20"},b=["textContent"],y=["textContent"],k={class:"w-full flex items-center justify-center"},v={__name:"error-404",props:{appName:{type:String,default:"Nuxt"},version:{type:String,default:""},statusCode:{type:String,default:"404"},statusMessage:{type:String,default:"Not Found"},description:{type:String,default:"Sorry, the page you are looking for could not be found."},backHome:{type:String,default:"Go back home"}},setup(t){const s=t;return r({title:`${s.statusCode} - ${s.statusMessage} | ${s.appName}`,script:[],style:[{children:""}]}),(S,w)=>{const o=u;return c(),d("div",f,[h,e("div",g,[e("h1",{class:"text-8xl sm:text-10xl font-medium mb-8",textContent:a(t.statusCode)},null,8,b),e("p",{class:"text-xl px-8 sm:px-0 sm:text-4xl font-light mb-16 leading-tight",textContent:a(t.description)},null,8,y),e("div",k,[i(o,{to:"/",class:"gradient-border text-md sm:text-xl py-2 px-4 sm:py-3 sm:px-6 cursor-pointer"},{default:l(()=>[p(a(t.backHome),1)]),_:1})])])])}}};var N=n(v,[["__scopeId","data-v-bed9aeba"]]);export{N as default};