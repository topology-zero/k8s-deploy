import{s as D,J as T,K as O,L as M,r as F,c as d,M as g,g as l,f as e,j as p,d as q,N as j,G as m,O as P,P as $,l as G,m as A,E as J,n as L,Q as K,R as Q,e as I,S as w,T as H}from"./index-cc90a63c.js";import{E as W}from"./el-pagination-1b5bba26.js";import{E as X,a as Y}from"./el-table-column-25359a9a.js";import{E as Z,a as ee}from"./el-select-ba1a1082.js";import{u as ae}from"./page-7fea1619.js";import{E as le,a as te}from"./el-radio-ca630d0f.js";import{v as oe}from"./directive-9d3e2d96.js";import{E as ne}from"./index-12cbd309.js";import{E as re}from"./index-2c9c9e7e.js";function se(n){return D({url:"/user",params:n})}function ue(n){return D({url:"/user/"+n,method:"get"})}function ie(n){return D({url:"/user",method:"post",data:n})}function de(n,_){return D({url:"/user/"+n,method:"put",data:_})}function me(n){return D({url:"/user/"+n,method:"delete"})}const pe={__name:"add-edit-dialog",props:{visible:{type:Boolean},visibleModifiers:{},formData:{type:Object},formDataModifiers:{}},emits:T(["done"],["update:visible","update:formData"]),setup(n,{emit:_}){const v=O(),{roleOptions:b}=v,E=_,f=M(n,"visible"),a=M(n,"formData"),V=F(),x=async()=>{await V.value.validate(),a.value.id==0?await ie(a.value):await de(a.value.id,a.value),f.value=!1,E("done"),$.success("成功")};return(z,t)=>{const y=G,c=A,h=ee,S=Z,s=le,r=te,C=J,u=L,k=K;return d(),g(k,{title:a.value.id==0?"添加用户":"编辑用户",modelValue:f.value,"onUpdate:modelValue":t[8]||(t[8]=o=>f.value=o),"close-on-click-modal":!1,width:"30%"},{footer:l(()=>[e(u,{onClick:t[6]||(t[6]=o=>f.value=!1)},{default:l(()=>[p("取 消")]),_:1}),e(u,{type:"primary",onClick:t[7]||(t[7]=o=>x())},{default:l(()=>[p("确 定")]),_:1})]),default:l(()=>[e(C,{ref_key:"elFormRef",ref:V,model:a.value,"label-width":"100px"},{default:l(()=>[e(c,{label:"用户名",rules:{required:!0,trigger:"blur",message:"请输入用户名"},prop:"username"},{default:l(()=>[e(y,{modelValue:a.value.username,"onUpdate:modelValue":t[0]||(t[0]=o=>a.value.username=o)},null,8,["modelValue"])]),_:1}),e(c,{label:"密码",rules:{required:a.value.id==0,trigger:"blur",message:"请输入密码"},prop:"password"},{default:l(()=>[e(y,{modelValue:a.value.password,"onUpdate:modelValue":t[1]||(t[1]=o=>a.value.password=o),placeholder:a.value.id==0?"":"不输入则不修改",type:"password"},null,8,["modelValue","placeholder"])]),_:1},8,["rules"]),e(c,{label:"所属角色",rules:{required:!0,trigger:"blur",message:"请选择角色"},prop:"roleId"},{default:l(()=>[e(S,{modelValue:a.value.roleId,"onUpdate:modelValue":t[2]||(t[2]=o=>a.value.roleId=o),style:{width:"100%"}},{default:l(()=>[(d(!0),q(P,null,j(m(b),(o,U)=>(d(),g(h,{key:U,label:o.name,value:o.id},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),e(c,{label:"真实姓名",rules:{required:!0,trigger:"blur",message:"请输入真实姓名"},prop:"realname"},{default:l(()=>[e(y,{modelValue:a.value.realname,"onUpdate:modelValue":t[3]||(t[3]=o=>a.value.realname=o)},null,8,["modelValue"])]),_:1}),e(c,{label:"电话号码",rules:{required:!0,trigger:"blur",max:12,message:"请输入合法的电话号码"},prop:"phone"},{default:l(()=>[e(y,{modelValue:a.value.phone,"onUpdate:modelValue":t[4]||(t[4]=o=>a.value.phone=o),oninput:"value=value.replace(/[^0-9.]/g,'')"},null,8,["modelValue"])]),_:1}),e(c,{label:"用户状态",rules:{required:!0,trigger:"blur",type:"enum",enum:[0,1],message:"请设置用户状态"},prop:"status"},{default:l(()=>[e(r,{modelValue:a.value.status,"onUpdate:modelValue":t[5]||(t[5]=o=>a.value.status=o)},{default:l(()=>[e(s,{label:0,border:""},{default:l(()=>[p("封停")]),_:1}),e(s,{label:1,border:""},{default:l(()=>[p("正常")]),_:1})]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1},8,["title","modelValue"])}}},ce={class:"app-container"},ge=I("p",{class:"page-title"},"用户列表",-1),fe={class:"filter-container"},_e={class:"page-container"},xe={__name:"index",setup(n){const _=()=>({id:0,username:"",password:"",realname:"",roleId:"",status:1,phone:""}),v=F(!1),b=F(_()),E=async()=>{const{data:s}=await se(V.value);a.value=s.data,z.value=s.total,f.value=!1},{tableLoading:f,list:a,params:V,pageSizes:x,total:z,handleSizeChange:t,handleCurrentChange:y}=ae(E),c=()=>{v.value=!0,b.value=_()},h=async s=>{v.value=!0;const{data:r}=await ue(s.id);b.value=r},S=async s=>{await ne.confirm("删除用户不可恢复","警告");const{message:r}=await me(s.id);$.success(r),E()};return(s,r)=>{const C=L,u=Y,k=re,o=X,U=W,B=Q("permission"),N=oe;return d(),q("div",ce,[ge,I("div",fe,[w((d(),g(C,{type:"primary",size:"default",plain:"",onClick:r[0]||(r[0]=i=>c())},{default:l(()=>[p(" 添加用户 ")]),_:1})),[[B,"admin:user:add"]])]),w((d(),g(o,{"element-loading-text":"加载中...",border:"",data:m(a)},{default:l(()=>[e(u,{label:"ID",prop:"id",align:"center",width:"60px"}),e(u,{align:"center",label:"用户名",prop:"username"}),e(u,{align:"center",label:"手机号",prop:"phone"}),e(u,{align:"center",label:"真实姓名",prop:"realname"}),e(u,{align:"center",label:"所属角色",prop:"rolename"}),e(u,{align:"center",label:"状态"},{default:l(({row:i})=>[i.status==1?(d(),g(k,{key:0,type:"success"},{default:l(()=>[p("正常")]),_:1})):(d(),g(k,{key:1,type:"danger"},{default:l(()=>[p("停封")]),_:1}))]),_:1}),e(u,{align:"center",width:"170px",label:"操作"},{default:l(({row:i})=>[w((d(),g(C,{size:"small",type:"primary",onClick:R=>h(i)},{default:l(()=>[p("编辑 ")]),_:2},1032,["onClick"])),[[B,"admin:user:edit"]]),i.id!=1?w((d(),g(C,{key:0,size:"small",type:"danger",onClick:R=>S(i)},{default:l(()=>[p(" 删除 ")]),_:2},1032,["onClick"])),[[B,"admin:user:del"]]):H("",!0)]),_:1})]),_:1},8,["data"])),[[N,m(f)]]),I("div",_e,[e(U,{background:"","current-page":m(V).page,"page-sizes":m(x),"page-size":m(V).pageSize,layout:"total, sizes, prev, pager, next, jumper",total:m(z),onSizeChange:m(t),onCurrentChange:m(y)},null,8,["current-page","page-sizes","page-size","total","onSizeChange","onCurrentChange"])]),e(pe,{visible:v.value,"onUpdate:visible":r[1]||(r[1]=i=>v.value=i),formData:b.value,"onUpdate:formData":r[2]||(r[2]=i=>b.value=i),onDone:E},null,8,["visible","formData"])])}}};export{xe as default};
