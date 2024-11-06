import{x as he,z as i,b1 as z,W as Y,aB as il,b2 as U,w as j,A as Te,Y as ge,t as se,Z as ne,aw as rl,G as we,V as ae,y as $e,C as et,H as F,S as te,a8 as ke,c as m,d as I,$ as k,e as $,I as B,aa as d,i as L,r as V,o as Me,b3 as G,T,ab as me,ai as lt,b4 as tt,b5 as nt,ax as at,ay as st,b6 as ot,b7 as P,b8 as it,ar as rt,aZ as ve,b9 as xe,az as ut,aW as dt,ao as ct,B as pt,ac as le,ba as Ie,ag as ft,an as Q,bb as vt,aq as ul,ad as bt,al as mt,am as ht,aC as be,bc as gt,ah as el,bd as yt,be as St,bf as Ct,bg as Ot,av as wt,bh as It,bi as Vt,a4 as Et,bj as Tt,a0 as dl,a7 as A,R as $t,f as H,g as R,O as ll,N as tl,j as nl,M as W,h as ee,bk as kt,a9 as Ve,bl as Mt,bm as Rt,ak as Dt,aA as cl}from"./index-cc90a63c.js";import{t as al,E as Lt}from"./index-2c9c9e7e.js";import{t as Bt,c as Nt,i as Ee,d as Wt}from"./el-table-column-25359a9a.js";var sl=1/0,Ft=17976931348623157e292;function zt(e){if(!e)return e===0?e:0;if(e=Bt(e),e===sl||e===-sl){var s=e<0?-1:1;return s*Ft}return e===e?e:0}function Pt(e){var s=zt(e),a=s%1;return s===s?a?s-a:s:0}function Kt(e,s,a,f){for(var o=e.length,g=a+(f?1:-1);f?g--:++g<o;)if(s(e[g],g,e))return g;return-1}var At=Math.max,Ht=Math.min;function Gt(e,s,a){var f=e==null?0:e.length;if(!f)return-1;var o=f-1;return a!==void 0&&(o=Pt(a),o=a<0?At(f+o,0):Ht(o,f-1)),Kt(e,Nt(s),o,!0)}const Ut=(e="")=>e.replace(/[|\\{}()[\]^$+*?.]/g,"\\$&").replace(/-/g,"\\x2d"),pl=Symbol("ElSelectGroup"),ye=Symbol("ElSelect");function jt(e,s){const a=he(ye),f=he(pl,{disabled:!1}),o=i(()=>y(z(a.props.modelValue),e.value)),g=i(()=>{var p;if(a.props.multiple){const b=z((p=a.props.modelValue)!=null?p:[]);return!o.value&&b.length>=a.props.multipleLimit&&a.props.multipleLimit>0}else return!1}),t=i(()=>e.label||(Y(e.value)?"":e.value)),v=i(()=>e.value||e.label||""),c=i(()=>e.disabled||s.groupDisabled||g.value),S=Te(),y=(p=[],b)=>{if(Y(e.value)){const u=a.props.valueKey;return p&&p.some(E=>il(U(E,u))===U(b,u))}else return p&&p.includes(b)},C=()=>{!e.disabled&&!f.disabled&&(a.states.hoveringIndex=a.optionsArray.indexOf(S.proxy))},w=p=>{const b=new RegExp(Ut(p),"i");s.visible=b.test(t.value)||e.created};return j(()=>t.value,()=>{!e.created&&!a.props.remote&&a.setSelected()}),j(()=>e.value,(p,b)=>{const{remote:u,valueKey:E}=a.props;if(p!==b&&(a.onOptionDestroy(b,S.proxy),a.onOptionCreate(S.proxy)),!e.created&&!u){if(E&&Y(p)&&Y(b)&&p[E]===b[E])return;a.setSelected()}}),j(()=>f.disabled,()=>{s.groupDisabled=f.disabled},{immediate:!0}),{select:a,currentLabel:t,currentValue:v,itemSelected:o,isDisabled:c,hoverItem:C,updateOption:w}}const qt=se({name:"ElOption",componentName:"ElOption",props:{value:{required:!0,type:[String,Number,Boolean,Object]},label:[String,Number],created:Boolean,disabled:Boolean},setup(e){const s=ne("select"),a=rl(),f=i(()=>[s.be("dropdown","item"),s.is("disabled",we(v)),s.is("selected",we(t)),s.is("hovering",we(w))]),o=ae({index:-1,groupDisabled:!1,visible:!0,hover:!1}),{currentLabel:g,itemSelected:t,isDisabled:v,select:c,hoverItem:S,updateOption:y}=jt(e,o),{visible:C,hover:w}=$e(o),p=Te().proxy;c.onOptionCreate(p),et(()=>{const u=p.value,{selected:E}=c.states,oe=(c.props.multiple?E:[E]).some(ie=>ie.value===p.value);F(()=>{c.states.cachedOptions.get(u)===p&&!oe&&c.states.cachedOptions.delete(u)}),c.onOptionDestroy(u,p)});function b(){v.value||c.handleOptionSelect(p)}return{ns:s,id:a,containerKls:f,currentLabel:g,itemSelected:t,isDisabled:v,select:c,hoverItem:S,updateOption:y,visible:C,hover:w,selectOptionClick:b,states:o}}});function Qt(e,s,a,f,o,g){return te((m(),I("li",{id:e.id,class:d(e.containerKls),role:"option","aria-disabled":e.isDisabled||void 0,"aria-selected":e.itemSelected,onMouseenter:e.hoverItem,onClick:L(e.selectOptionClick,["stop"])},[k(e.$slots,"default",{},()=>[$("span",null,B(e.currentLabel),1)])],42,["id","aria-disabled","aria-selected","onMouseenter","onClick"])),[[ke,e.visible]])}var Re=ge(qt,[["render",Qt],["__file","option.vue"]]);const Yt=se({name:"ElSelectDropdown",componentName:"ElSelectDropdown",setup(){const e=he(ye),s=ne("select"),a=i(()=>e.props.popperClass),f=i(()=>e.props.multiple),o=i(()=>e.props.fitInputWidth),g=V("");function t(){var v;g.value=`${(v=e.selectRef)==null?void 0:v.offsetWidth}px`}return Me(()=>{t(),G(e.selectRef,t)}),{ns:s,minWidth:g,popperClass:a,isMultiple:f,isFitInputWidth:o}}});function Zt(e,s,a,f,o,g){return m(),I("div",{class:d([e.ns.b("dropdown"),e.ns.is("multiple",e.isMultiple),e.popperClass]),style:me({[e.isFitInputWidth?"width":"minWidth"]:e.minWidth})},[e.$slots.header?(m(),I("div",{key:0,class:d(e.ns.be("dropdown","header"))},[k(e.$slots,"header")],2)):T("v-if",!0),k(e.$slots,"default"),e.$slots.footer?(m(),I("div",{key:1,class:d(e.ns.be("dropdown","footer"))},[k(e.$slots,"footer")],2)):T("v-if",!0)],6)}var Xt=ge(Yt,[["render",Zt],["__file","select-dropdown.vue"]]);const Jt=11,_t=(e,s)=>{const{t:a}=lt(),f=rl(),o=ne("select"),g=ne("input"),t=ae({inputValue:"",options:new Map,cachedOptions:new Map,disabledOptions:new Map,optionValues:[],selected:[],selectionWidth:0,calculatorWidth:0,collapseItemWidth:0,selectedLabel:"",hoveringIndex:-1,previousQuery:null,inputHovering:!1,menuVisibleOnFocus:!1,isBeforeHide:!1}),v=V(null),c=V(null),S=V(null),y=V(null),C=V(null),w=V(null),p=V(null),b=V(null),u=V(null),E=V(null),Z=V(null),oe=V(null),{isComposing:ie,handleCompositionStart:vl,handleCompositionUpdate:bl,handleCompositionEnd:ml}=tt({afterComposition:l=>je(l)}),{wrapperRef:De,isFocused:Le,handleBlur:hl}=nt(C,{beforeFocus(){return _.value},afterFocus(){e.automaticDropdown&&!O.value&&(O.value=!0,t.menuVisibleOnFocus=!0)},beforeBlur(l){var n,r;return((n=S.value)==null?void 0:n.isFocusInsideContent(l))||((r=y.value)==null?void 0:r.isFocusInsideContent(l))},afterBlur(){O.value=!1,t.menuVisibleOnFocus=!1}}),O=V(!1),X=V(),{form:Be,formItem:J}=at(),{inputId:gl}=st(e,{formItemContext:J}),{valueOnClear:yl,isEmptyValue:Sl}=ot(e),_=i(()=>e.disabled||(Be==null?void 0:Be.disabled)),Se=i(()=>P(e.modelValue)?e.modelValue.length>0:!Sl(e.modelValue)),Cl=i(()=>e.clearable&&!_.value&&t.inputHovering&&Se.value),Ne=i(()=>e.remote&&e.filterable&&!e.remoteShowSuffix?"":e.suffixIcon),Ol=i(()=>o.is("reverse",Ne.value&&O.value)),We=i(()=>(J==null?void 0:J.validateState)||""),wl=i(()=>it[We.value]),Il=i(()=>e.remote?300:0),Fe=i(()=>e.loading?e.loadingText||a("el.select.loading"):e.remote&&!t.inputValue&&t.options.size===0?!1:e.filterable&&t.inputValue&&t.options.size>0&&re.value===0?e.noMatchText||a("el.select.noMatch"):t.options.size===0?e.noDataText||a("el.select.noData"):null),re=i(()=>M.value.filter(l=>l.visible).length),M=i(()=>{const l=Array.from(t.options.values()),n=[];return t.optionValues.forEach(r=>{const h=l.findIndex(N=>N.value===r);h>-1&&n.push(l[h])}),n.length>=l.length?n:l}),Vl=i(()=>Array.from(t.cachedOptions.values())),El=i(()=>{const l=M.value.filter(n=>!n.created).some(n=>n.currentLabel===t.inputValue);return e.filterable&&e.allowCreate&&t.inputValue!==""&&!l}),ze=()=>{e.filterable&&le(e.filterMethod)||e.filterable&&e.remote&&le(e.remoteMethod)||M.value.forEach(l=>{var n;(n=l.updateOption)==null||n.call(l,t.inputValue)})},Pe=rt(),Tl=i(()=>["small"].includes(Pe.value)?"small":"default"),$l=i({get(){return O.value&&Fe.value!==!1},set(l){O.value=l}}),kl=i(()=>{if(e.multiple&&!ve(e.modelValue))return z(e.modelValue).length===0&&!t.inputValue;const l=P(e.modelValue)?e.modelValue[0]:e.modelValue;return e.filterable||ve(l)?!t.inputValue:!0}),Ml=i(()=>{var l;const n=(l=e.placeholder)!=null?l:a("el.select.placeholder");return e.multiple||!Se.value?n:t.selectedLabel}),Rl=i(()=>xe?null:"mouseenter");j(()=>e.modelValue,(l,n)=>{e.multiple&&e.filterable&&!e.reserveKeyword&&(t.inputValue="",ue("")),de(),!Ee(l,n)&&e.validateEvent&&(J==null||J.validate("change").catch(r=>ut()))},{flush:"post",deep:!0}),j(()=>O.value,l=>{l?ue(t.inputValue):(t.inputValue="",t.previousQuery=null,t.isBeforeHide=!0),s("visible-change",l)}),j(()=>t.options.entries(),()=>{var l;if(!dt)return;const n=((l=v.value)==null?void 0:l.querySelectorAll("input"))||[];(!e.filterable&&!e.defaultFirstOption&&!ve(e.modelValue)||!Array.from(n).includes(document.activeElement))&&de(),e.defaultFirstOption&&(e.filterable||e.remote)&&re.value&&Ke()},{flush:"post"}),j(()=>t.hoveringIndex,l=>{ct(l)&&l>-1?X.value=M.value[l]||{}:X.value={},M.value.forEach(n=>{n.hover=X.value===n})}),pt(()=>{t.isBeforeHide||ze()});const ue=l=>{t.previousQuery===l||ie.value||(t.previousQuery=l,e.filterable&&le(e.filterMethod)?e.filterMethod(l):e.filterable&&e.remote&&le(e.remoteMethod)&&e.remoteMethod(l),e.defaultFirstOption&&(e.filterable||e.remote)&&re.value?F(Ke):F(Dl))},Ke=()=>{const l=M.value.filter(h=>h.visible&&!h.disabled&&!h.states.groupDisabled),n=l.find(h=>h.created),r=l[0];t.hoveringIndex=Ze(M.value,n||r)},de=()=>{if(e.multiple)t.selectedLabel="";else{const n=P(e.modelValue)?e.modelValue[0]:e.modelValue,r=Ae(n);t.selectedLabel=r.currentLabel,t.selected=[r];return}const l=[];ve(e.modelValue)||z(e.modelValue).forEach(n=>{l.push(Ae(n))}),t.selected=l},Ae=l=>{let n;const r=Ie(l).toLowerCase()==="object",h=Ie(l).toLowerCase()==="null",N=Ie(l).toLowerCase()==="undefined";for(let K=t.cachedOptions.size-1;K>=0;K--){const D=Vl.value[K];if(r?U(D.value,e.valueKey)===U(l,e.valueKey):D.value===l){n={value:l,currentLabel:D.currentLabel,get isDisabled(){return D.isDisabled}};break}}if(n)return n;const q=r?l.label:!h&&!N?l:"";return{value:l,currentLabel:q}},Dl=()=>{t.hoveringIndex=M.value.findIndex(l=>t.selected.some(n=>Oe(n)===Oe(l)))},Ll=()=>{t.selectionWidth=c.value.getBoundingClientRect().width},He=()=>{t.calculatorWidth=w.value.getBoundingClientRect().width},Bl=()=>{t.collapseItemWidth=Z.value.getBoundingClientRect().width},Ce=()=>{var l,n;(n=(l=S.value)==null?void 0:l.updatePopper)==null||n.call(l)},Ge=()=>{var l,n;(n=(l=y.value)==null?void 0:l.updatePopper)==null||n.call(l)},Ue=()=>{t.inputValue.length>0&&!O.value&&(O.value=!0),ue(t.inputValue)},je=l=>{if(t.inputValue=l.target.value,e.remote)qe();else return Ue()},qe=Wt(()=>{Ue()},Il.value),x=l=>{Ee(e.modelValue,l)||s(ul,l)},Nl=l=>Gt(l,n=>!t.disabledOptions.has(n)),Wl=l=>{if(e.multiple&&l.code!==ft.delete&&l.target.value.length<=0){const n=z(e.modelValue).slice(),r=Nl(n);if(r<0)return;const h=n[r];n.splice(r,1),s(Q,n),x(n),s("remove-tag",h)}},Fl=(l,n)=>{const r=t.selected.indexOf(n);if(r>-1&&!_.value){const h=z(e.modelValue).slice();h.splice(r,1),s(Q,h),x(h),s("remove-tag",n.value)}l.stopPropagation(),pe()},Qe=l=>{l.stopPropagation();const n=e.multiple?[]:yl.value;if(e.multiple)for(const r of t.selected)r.isDisabled&&n.push(r.value);s(Q,n),x(n),t.hoveringIndex=-1,O.value=!1,s("clear"),pe()},Ye=l=>{var n;if(e.multiple){const r=z((n=e.modelValue)!=null?n:[]).slice(),h=Ze(r,l.value);h>-1?r.splice(h,1):(e.multipleLimit<=0||r.length<e.multipleLimit)&&r.push(l.value),s(Q,r),x(r),l.created&&ue(""),e.filterable&&!e.reserveKeyword&&(t.inputValue="")}else s(Q,l.value),x(l.value),O.value=!1;pe(),!O.value&&F(()=>{ce(l)})},Ze=(l=[],n)=>{if(!Y(n))return l.indexOf(n);const r=e.valueKey;let h=-1;return l.some((N,q)=>il(U(N,r))===U(n,r)?(h=q,!0):!1),h},ce=l=>{var n,r,h,N,q;const fe=P(l)?l[0]:l;let K=null;if(fe!=null&&fe.value){const D=M.value.filter(_e=>_e.value===fe.value);D.length>0&&(K=D[0].$el)}if(S.value&&K){const D=(N=(h=(r=(n=S.value)==null?void 0:n.popperRef)==null?void 0:r.contentRef)==null?void 0:h.querySelector)==null?void 0:N.call(h,`.${o.be("dropdown","wrap")}`);D&&vt(D,K)}(q=oe.value)==null||q.handleScroll()},zl=l=>{t.options.set(l.value,l),t.cachedOptions.set(l.value,l),l.disabled&&t.disabledOptions.set(l.value,l)},Pl=(l,n)=>{t.options.get(l)===n&&t.options.delete(l)},Kl=i(()=>{var l,n;return(n=(l=S.value)==null?void 0:l.popperRef)==null?void 0:n.contentRef}),Al=()=>{t.isBeforeHide=!1,F(()=>ce(t.selected))},pe=()=>{var l;(l=C.value)==null||l.focus()},Hl=()=>{var l;if(O.value){O.value=!1,F(()=>{var n;return(n=C.value)==null?void 0:n.blur()});return}(l=C.value)==null||l.blur()},Gl=l=>{Qe(l)},Ul=l=>{if(O.value=!1,Le.value){const n=new FocusEvent("focus",l);F(()=>hl(n))}},jl=()=>{t.inputValue.length>0?t.inputValue="":O.value=!1},Xe=()=>{_.value||(xe&&(t.inputHovering=!0),t.menuVisibleOnFocus?t.menuVisibleOnFocus=!1:O.value=!O.value)},ql=()=>{O.value?M.value[t.hoveringIndex]&&Ye(M.value[t.hoveringIndex]):Xe()},Oe=l=>Y(l.value)?U(l.value,e.valueKey):l.value,Ql=i(()=>M.value.filter(l=>l.visible).every(l=>l.disabled)),Yl=i(()=>e.multiple?e.collapseTags?t.selected.slice(0,e.maxCollapseTags):t.selected:[]),Zl=i(()=>e.multiple?e.collapseTags?t.selected.slice(e.maxCollapseTags):[]:[]),Je=l=>{if(!O.value){O.value=!0;return}if(!(t.options.size===0||t.filteredOptionsCount===0||ie.value)&&!Ql.value){l==="next"?(t.hoveringIndex++,t.hoveringIndex===t.options.size&&(t.hoveringIndex=0)):l==="prev"&&(t.hoveringIndex--,t.hoveringIndex<0&&(t.hoveringIndex=t.options.size-1));const n=M.value[t.hoveringIndex];(n.disabled===!0||n.states.groupDisabled===!0||!n.visible)&&Je(l),F(()=>ce(X.value))}},Xl=()=>{if(!c.value)return 0;const l=window.getComputedStyle(c.value);return Number.parseFloat(l.gap||"6px")},Jl=i(()=>{const l=Xl();return{maxWidth:`${Z.value&&e.maxCollapseTags===1?t.selectionWidth-t.collapseItemWidth-l:t.selectionWidth}px`}}),_l=i(()=>({maxWidth:`${t.selectionWidth}px`})),xl=i(()=>({width:`${Math.max(t.calculatorWidth,Jt)}px`}));return G(c,Ll),G(w,He),G(u,Ce),G(De,Ce),G(E,Ge),G(Z,Bl),Me(()=>{de()}),{inputId:gl,contentId:f,nsSelect:o,nsInput:g,states:t,isFocused:Le,expanded:O,optionsArray:M,hoverOption:X,selectSize:Pe,filteredOptionsCount:re,resetCalculatorWidth:He,updateTooltip:Ce,updateTagTooltip:Ge,debouncedOnInputChange:qe,onInput:je,deletePrevTag:Wl,deleteTag:Fl,deleteSelected:Qe,handleOptionSelect:Ye,scrollToOption:ce,hasModelValue:Se,shouldShowPlaceholder:kl,currentPlaceholder:Ml,mouseEnterEventName:Rl,showClose:Cl,iconComponent:Ne,iconReverse:Ol,validateState:We,validateIcon:wl,showNewOption:El,updateOptions:ze,collapseTagSize:Tl,setSelected:de,selectDisabled:_,emptyText:Fe,handleCompositionStart:vl,handleCompositionUpdate:bl,handleCompositionEnd:ml,onOptionCreate:zl,onOptionDestroy:Pl,handleMenuEnter:Al,focus:pe,blur:Hl,handleClearClick:Gl,handleClickOutside:Ul,handleEsc:jl,toggleMenu:Xe,selectOption:ql,getValueKey:Oe,navigateOptions:Je,dropdownMenuVisible:$l,showTagList:Yl,collapseTagList:Zl,tagStyle:Jl,collapseTagStyle:_l,inputStyle:xl,popperRef:Kl,inputRef:C,tooltipRef:S,tagTooltipRef:y,calculatorRef:w,prefixRef:p,suffixRef:b,selectRef:v,wrapperRef:De,selectionRef:c,scrollbarRef:oe,menuRef:u,tagMenuRef:E,collapseItemRef:Z}};var xt=se({name:"ElOptions",setup(e,{slots:s}){const a=he(ye);let f=[];return()=>{var o,g;const t=(o=s.default)==null?void 0:o.call(s),v=[];function c(S){P(S)&&S.forEach(y=>{var C,w,p,b;const u=(C=(y==null?void 0:y.type)||{})==null?void 0:C.name;u==="ElOptionGroup"?c(!bt(y.children)&&!P(y.children)&&le((w=y.children)==null?void 0:w.default)?(p=y.children)==null?void 0:p.default():y.children):u==="ElOption"?v.push((b=y.props)==null?void 0:b.value):P(y.children)&&c(y.children)})}return t.length&&c((g=t[0])==null?void 0:g.children),Ee(v,f)||(f=v,a&&(a.states.optionValues=v)),t}}});const en=mt({name:String,id:String,modelValue:{type:[Array,String,Number,Boolean,Object],default:void 0},autocomplete:{type:String,default:"off"},automaticDropdown:Boolean,size:ht,effect:{type:be(String),default:"light"},disabled:Boolean,clearable:Boolean,filterable:Boolean,allowCreate:Boolean,loading:Boolean,popperClass:{type:String,default:""},popperOptions:{type:be(Object),default:()=>({})},remote:Boolean,loadingText:String,noMatchText:String,noDataText:String,remoteMethod:Function,filterMethod:Function,multiple:Boolean,multipleLimit:{type:Number,default:0},placeholder:{type:String},defaultFirstOption:Boolean,reserveKeyword:{type:Boolean,default:!0},valueKey:{type:String,default:"value"},collapseTags:Boolean,collapseTagsTooltip:Boolean,maxCollapseTags:{type:Number,default:1},teleported:gt.teleported,persistent:{type:Boolean,default:!0},clearIcon:{type:el,default:yt},fitInputWidth:Boolean,suffixIcon:{type:el,default:St},tagType:{...al.type,default:"info"},tagEffect:{...al.effect,default:"light"},validateEvent:{type:Boolean,default:!0},remoteShowSuffix:Boolean,placement:{type:be(String),values:Ct,default:"bottom-start"},fallbackPlacements:{type:be(Array),default:["bottom-start","top-start","right","left"]},appendTo:String,...Ot,...wt(["ariaLabel"])}),ol="ElSelect",ln=se({name:ol,componentName:ol,components:{ElSelectMenu:Xt,ElOption:Re,ElOptions:xt,ElTag:Lt,ElScrollbar:It,ElTooltip:Vt,ElIcon:Et},directives:{ClickOutside:Tt},props:en,emits:[Q,ul,"remove-tag","clear","visible-change","focus","blur"],setup(e,{emit:s}){const a=i(()=>{const{modelValue:t,multiple:v}=e,c=v?[]:void 0;return P(t)?v?t:c:v?c:t}),f=ae({...$e(e),modelValue:a}),o=_t(f,s);dl(ye,ae({props:f,states:o.states,optionsArray:o.optionsArray,handleOptionSelect:o.handleOptionSelect,onOptionCreate:o.onOptionCreate,onOptionDestroy:o.onOptionDestroy,selectRef:o.selectRef,setSelected:o.setSelected}));const g=i(()=>e.multiple?o.states.selected.map(t=>t.currentLabel):o.states.selectedLabel);return{...o,modelValue:a,selectedLabel:g}}});function tn(e,s,a,f,o,g){const t=A("el-tag"),v=A("el-tooltip"),c=A("el-icon"),S=A("el-option"),y=A("el-options"),C=A("el-scrollbar"),w=A("el-select-menu"),p=$t("click-outside");return te((m(),I("div",{ref:"selectRef",class:d([e.nsSelect.b(),e.nsSelect.m(e.selectSize)]),[Mt(e.mouseEnterEventName)]:b=>e.states.inputHovering=!0,onMouseleave:b=>e.states.inputHovering=!1},[H(v,{ref:"tooltipRef",visible:e.dropdownMenuVisible,placement:e.placement,teleported:e.teleported,"popper-class":[e.nsSelect.e("popper"),e.popperClass],"popper-options":e.popperOptions,"fallback-placements":e.fallbackPlacements,effect:e.effect,pure:"",trigger:"click",transition:`${e.nsSelect.namespace.value}-zoom-in-top`,"stop-popper-mouse-event":!1,"gpu-acceleration":!1,persistent:e.persistent,"append-to":e.appendTo,onBeforeShow:e.handleMenuEnter,onHide:b=>e.states.isBeforeHide=!1},{default:R(()=>{var b;return[$("div",{ref:"wrapperRef",class:d([e.nsSelect.e("wrapper"),e.nsSelect.is("focused",e.isFocused),e.nsSelect.is("hovering",e.states.inputHovering),e.nsSelect.is("filterable",e.filterable),e.nsSelect.is("disabled",e.selectDisabled)]),onClick:L(e.toggleMenu,["prevent"])},[e.$slots.prefix?(m(),I("div",{key:0,ref:"prefixRef",class:d(e.nsSelect.e("prefix"))},[k(e.$slots,"prefix")],2)):T("v-if",!0),$("div",{ref:"selectionRef",class:d([e.nsSelect.e("selection"),e.nsSelect.is("near",e.multiple&&!e.$slots.prefix&&!!e.states.selected.length)])},[e.multiple?k(e.$slots,"tag",{key:0},()=>[(m(!0),I(ll,null,tl(e.showTagList,u=>(m(),I("div",{key:e.getValueKey(u),class:d(e.nsSelect.e("selected-item"))},[H(t,{closable:!e.selectDisabled&&!u.isDisabled,size:e.collapseTagSize,type:e.tagType,effect:e.tagEffect,"disable-transitions":"",style:me(e.tagStyle),onClose:E=>e.deleteTag(E,u)},{default:R(()=>[$("span",{class:d(e.nsSelect.e("tags-text"))},[k(e.$slots,"label",{label:u.currentLabel,value:u.value},()=>[nl(B(u.currentLabel),1)])],2)]),_:2},1032,["closable","size","type","effect","style","onClose"])],2))),128)),e.collapseTags&&e.states.selected.length>e.maxCollapseTags?(m(),W(v,{key:0,ref:"tagTooltipRef",disabled:e.dropdownMenuVisible||!e.collapseTagsTooltip,"fallback-placements":["bottom","top","right","left"],effect:e.effect,placement:"bottom",teleported:e.teleported},{default:R(()=>[$("div",{ref:"collapseItemRef",class:d(e.nsSelect.e("selected-item"))},[H(t,{closable:!1,size:e.collapseTagSize,type:e.tagType,effect:e.tagEffect,"disable-transitions":"",style:me(e.collapseTagStyle)},{default:R(()=>[$("span",{class:d(e.nsSelect.e("tags-text"))}," + "+B(e.states.selected.length-e.maxCollapseTags),3)]),_:1},8,["size","type","effect","style"])],2)]),content:R(()=>[$("div",{ref:"tagMenuRef",class:d(e.nsSelect.e("selection"))},[(m(!0),I(ll,null,tl(e.collapseTagList,u=>(m(),I("div",{key:e.getValueKey(u),class:d(e.nsSelect.e("selected-item"))},[H(t,{class:"in-tooltip",closable:!e.selectDisabled&&!u.isDisabled,size:e.collapseTagSize,type:e.tagType,effect:e.tagEffect,"disable-transitions":"",onClose:E=>e.deleteTag(E,u)},{default:R(()=>[$("span",{class:d(e.nsSelect.e("tags-text"))},[k(e.$slots,"label",{label:u.currentLabel,value:u.value},()=>[nl(B(u.currentLabel),1)])],2)]),_:2},1032,["closable","size","type","effect","onClose"])],2))),128))],2)]),_:3},8,["disabled","effect","teleported"])):T("v-if",!0)]):T("v-if",!0),e.selectDisabled?T("v-if",!0):(m(),I("div",{key:1,class:d([e.nsSelect.e("selected-item"),e.nsSelect.e("input-wrapper"),e.nsSelect.is("hidden",!e.filterable)])},[te($("input",{id:e.inputId,ref:"inputRef","onUpdate:modelValue":u=>e.states.inputValue=u,type:"text",name:e.name,class:d([e.nsSelect.e("input"),e.nsSelect.is(e.selectSize)]),disabled:e.selectDisabled,autocomplete:e.autocomplete,style:me(e.inputStyle),role:"combobox",readonly:!e.filterable,spellcheck:"false","aria-activedescendant":((b=e.hoverOption)==null?void 0:b.id)||"","aria-controls":e.contentId,"aria-expanded":e.dropdownMenuVisible,"aria-label":e.ariaLabel,"aria-autocomplete":"none","aria-haspopup":"listbox",onKeydown:[ee(L(u=>e.navigateOptions("next"),["stop","prevent"]),["down"]),ee(L(u=>e.navigateOptions("prev"),["stop","prevent"]),["up"]),ee(L(e.handleEsc,["stop","prevent"]),["esc"]),ee(L(e.selectOption,["stop","prevent"]),["enter"]),ee(L(e.deletePrevTag,["stop"]),["delete"])],onCompositionstart:e.handleCompositionStart,onCompositionupdate:e.handleCompositionUpdate,onCompositionend:e.handleCompositionEnd,onInput:e.onInput,onClick:L(e.toggleMenu,["stop"])},null,46,["id","onUpdate:modelValue","name","disabled","autocomplete","readonly","aria-activedescendant","aria-controls","aria-expanded","aria-label","onKeydown","onCompositionstart","onCompositionupdate","onCompositionend","onInput","onClick"]),[[kt,e.states.inputValue]]),e.filterable?(m(),I("span",{key:0,ref:"calculatorRef","aria-hidden":"true",class:d(e.nsSelect.e("input-calculator")),textContent:B(e.states.inputValue)},null,10,["textContent"])):T("v-if",!0)],2)),e.shouldShowPlaceholder?(m(),I("div",{key:2,class:d([e.nsSelect.e("selected-item"),e.nsSelect.e("placeholder"),e.nsSelect.is("transparent",!e.hasModelValue||e.expanded&&!e.states.inputValue)])},[e.hasModelValue?k(e.$slots,"label",{key:0,label:e.currentPlaceholder,value:e.modelValue},()=>[$("span",null,B(e.currentPlaceholder),1)]):(m(),I("span",{key:1},B(e.currentPlaceholder),1))],2)):T("v-if",!0)],2),$("div",{ref:"suffixRef",class:d(e.nsSelect.e("suffix"))},[e.iconComponent&&!e.showClose?(m(),W(c,{key:0,class:d([e.nsSelect.e("caret"),e.nsSelect.e("icon"),e.iconReverse])},{default:R(()=>[(m(),W(Ve(e.iconComponent)))]),_:1},8,["class"])):T("v-if",!0),e.showClose&&e.clearIcon?(m(),W(c,{key:1,class:d([e.nsSelect.e("caret"),e.nsSelect.e("icon"),e.nsSelect.e("clear")]),onClick:e.handleClearClick},{default:R(()=>[(m(),W(Ve(e.clearIcon)))]),_:1},8,["class","onClick"])):T("v-if",!0),e.validateState&&e.validateIcon?(m(),W(c,{key:2,class:d([e.nsInput.e("icon"),e.nsInput.e("validateIcon")])},{default:R(()=>[(m(),W(Ve(e.validateIcon)))]),_:1},8,["class"])):T("v-if",!0)],2)],10,["onClick"])]}),content:R(()=>[H(w,{ref:"menuRef"},{default:R(()=>[e.$slots.header?(m(),I("div",{key:0,class:d(e.nsSelect.be("dropdown","header")),onClick:L(()=>{},["stop"])},[k(e.$slots,"header")],10,["onClick"])):T("v-if",!0),te(H(C,{id:e.contentId,ref:"scrollbarRef",tag:"ul","wrap-class":e.nsSelect.be("dropdown","wrap"),"view-class":e.nsSelect.be("dropdown","list"),class:d([e.nsSelect.is("empty",e.filteredOptionsCount===0)]),role:"listbox","aria-label":e.ariaLabel,"aria-orientation":"vertical"},{default:R(()=>[e.showNewOption?(m(),W(S,{key:0,value:e.states.inputValue,created:!0},null,8,["value"])):T("v-if",!0),H(y,null,{default:R(()=>[k(e.$slots,"default")]),_:3})]),_:3},8,["id","wrap-class","view-class","class","aria-label"]),[[ke,e.states.options.size>0&&!e.loading]]),e.$slots.loading&&e.loading?(m(),I("div",{key:1,class:d(e.nsSelect.be("dropdown","loading"))},[k(e.$slots,"loading")],2)):e.loading||e.filteredOptionsCount===0?(m(),I("div",{key:2,class:d(e.nsSelect.be("dropdown","empty"))},[k(e.$slots,"empty",{},()=>[$("span",null,B(e.emptyText),1)])],2)):T("v-if",!0),e.$slots.footer?(m(),I("div",{key:3,class:d(e.nsSelect.be("dropdown","footer")),onClick:L(()=>{},["stop"])},[k(e.$slots,"footer")],10,["onClick"])):T("v-if",!0)]),_:3},512)]),_:3},8,["visible","placement","teleported","popper-class","popper-options","fallback-placements","effect","transition","persistent","append-to","onBeforeShow","onHide"])],16,["onMouseleave"])),[[p,e.handleClickOutside,e.popperRef]])}var nn=ge(ln,[["render",tn],["__file","select.vue"]]);const an=se({name:"ElOptionGroup",componentName:"ElOptionGroup",props:{label:String,disabled:Boolean},setup(e){const s=ne("select"),a=V(null),f=Te(),o=V([]);dl(pl,ae({...$e(e)}));const g=i(()=>o.value.some(S=>S.visible===!0)),t=S=>{var y,C;return((y=S.type)==null?void 0:y.name)==="ElOption"&&!!((C=S.component)!=null&&C.proxy)},v=S=>{const y=z(S),C=[];return y.forEach(w=>{var p,b;t(w)?C.push(w.component.proxy):(p=w.children)!=null&&p.length?C.push(...v(w.children)):(b=w.component)!=null&&b.subTree&&C.push(...v(w.component.subTree))}),C},c=()=>{o.value=v(f.subTree)};return Me(()=>{c()}),Rt(a,c,{attributes:!0,subtree:!0,childList:!0}),{groupRef:a,visible:g,ns:s}}});function sn(e,s,a,f,o,g){return te((m(),I("ul",{ref:"groupRef",class:d(e.ns.be("group","wrap"))},[$("li",{class:d(e.ns.be("group","title"))},B(e.label),3),$("li",null,[$("ul",{class:d(e.ns.b("group"))},[k(e.$slots,"default")],2)])],2)),[[ke,e.visible]])}var fl=ge(an,[["render",sn],["__file","option-group.vue"]]);const dn=Dt(nn,{Option:Re,OptionGroup:fl}),cn=cl(Re);cl(fl);export{dn as E,cn as a,ye as s};
