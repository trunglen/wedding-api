webpackJsonp([5],{Vtsp:function(l,n,u){"use strict";Object.defineProperty(n,"__esModule",{value:!0});var e=u("WT6e"),t=function(){},d=u("7DMc"),o=u("Xjw4"),i=u("AjD5"),a=u("typQ"),s=u("bfOx"),r=function(){function l(l,n,u,e){this.userService=l,this.notifycationService=n,this.router=u,this.notificationService=e}return l.prototype.ngOnInit=function(){this.users$=this.userService.getUsers("manager")},l.prototype.onCreateManager=function(l){var n=this,u=l.value;u.role="manager",this.userService.createUser(u).subscribe(function(l){n.notifycationService.success("T\u1ea1o ng\u01b0\u1eddi qu\u1ea3n l\xed th\xe0nh c\xf4ng")})},l.prototype.onEdit=function(l){this.router.navigate(["update/"+l.id],{queryParams:l})},l.prototype.onDelete=function(l){var n=this;this.notificationService.confirm("B\u1ea1n c\xf3 ch\u1eafc mu\u1ed1n x\xf3a").subscribe(function(u){n.userService.deleteUser(l).subscribe(function(l){n.notificationService.success("\u0110\xe3 x\xf3a th\xe0nh c\xf4ng")})})},l}(),c=e["\u0275crt"]({encapsulation:0,styles:[[""]],data:{}});function g(l){return e["\u0275vid"](0,[(l()(),e["\u0275eld"](0,0,null,null,26,"tr",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](2,0,null,null,1,"td",[],null,null,null,null,null)),(l()(),e["\u0275ted"](3,null,["",""])),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](5,0,null,null,1,"td",[],null,null,null,null,null)),(l()(),e["\u0275ted"](6,null,["",""])),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](8,0,null,null,4,"td",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n                "])),(l()(),e["\u0275eld"](10,0,null,null,1,"span",[["class","badge badge-success"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["Active"])),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](14,0,null,null,11,"td",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n                  "])),(l()(),e["\u0275eld"](16,0,null,null,3,"button",[["class","btn btn-sm btn-danger"],["type","button"]],null,[[null,"click"]],function(l,n,u){var e=!0;return"click"===n&&(e=!1!==l.component.onEdit(l.context.$implicit)&&e),e},null,null)),(l()(),e["\u0275ted"](-1,null,["\n                    "])),(l()(),e["\u0275eld"](18,0,null,null,0,"i",[["class","fa fa-edit"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,[" S\u1eeda"])),(l()(),e["\u0275ted"](-1,null,["\n                    "])),(l()(),e["\u0275eld"](21,0,null,null,3,"button",[["class","btn btn-sm btn-default"],["type","button"]],null,[[null,"click"]],function(l,n,u){var e=!0;return"click"===n&&(e=!1!==l.component.onDelete(l.context.$implicit.id)&&e),e},null,null)),(l()(),e["\u0275ted"](-1,null,["\n                      "])),(l()(),e["\u0275eld"](23,0,null,null,0,"i",[["class","fa fa-trash"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,[" X\xf3a"])),(l()(),e["\u0275ted"](-1,null,["\n                "])),(l()(),e["\u0275ted"](-1,null,["\n            "]))],null,function(l,n){l(n,3,0,n.context.$implicit.name),l(n,6,0,n.context.$implicit.phone)})}function p(l){return e["\u0275vid"](0,[(l()(),e["\u0275eld"](0,0,null,null,94,"div",[["class","row"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n  "])),(l()(),e["\u0275eld"](2,0,null,null,91,"div",[["class","col-lg-12"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n    "])),(l()(),e["\u0275eld"](4,0,null,null,88,"form",[["action",""],["novalidate",""]],[[2,"ng-untouched",null],[2,"ng-touched",null],[2,"ng-pristine",null],[2,"ng-dirty",null],[2,"ng-valid",null],[2,"ng-invalid",null],[2,"ng-pending",null]],[[null,"submit"],[null,"reset"]],function(l,n,u){var t=!0;return"submit"===n&&(t=!1!==e["\u0275nov"](l,6).onSubmit(u)&&t),"reset"===n&&(t=!1!==e["\u0275nov"](l,6).onReset()&&t),t},null,null)),e["\u0275did"](5,16384,null,0,d["\u0275bf"],[],null,null),e["\u0275did"](6,4210688,[["f",4]],0,d.NgForm,[[8,null],[8,null]],null,null),e["\u0275prd"](2048,null,d.ControlContainer,null,[d.NgForm]),e["\u0275did"](8,16384,null,0,d.NgControlStatusGroup,[d.ControlContainer],null,null),(l()(),e["\u0275ted"](-1,null,["\n    "])),(l()(),e["\u0275eld"](10,0,null,null,81,"div",[["class","card"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275eld"](12,0,null,null,4,"div",[["class","card-header"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275eld"](14,0,null,null,1,"strong",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["T\u1ea1o account qu\u1ea3n l\xed"])),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275eld"](18,0,null,null,59,"div",[["class","card-body"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275eld"](20,0,null,null,17,"div",[["class","row"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275eld"](22,0,null,null,14,"div",[["class","col-sm-12"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](24,0,null,null,11,"div",[["class","form-group"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n                "])),(l()(),e["\u0275eld"](26,0,null,null,1,"label",[["for","name"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["T\xean ng\u01b0\u1eddi qu\u1ea3n l\xed"])),(l()(),e["\u0275ted"](-1,null,["\n                "])),(l()(),e["\u0275eld"](29,0,null,null,5,"input",[["class","form-control"],["id","name"],["name","name"],["ngModel",""],["placeholder","Nh\u1eadp t\xean ng\u01b0\u1eddi qu\u1ea3n l\xed"],["type","text"]],[[2,"ng-untouched",null],[2,"ng-touched",null],[2,"ng-pristine",null],[2,"ng-dirty",null],[2,"ng-valid",null],[2,"ng-invalid",null],[2,"ng-pending",null]],[[null,"input"],[null,"blur"],[null,"compositionstart"],[null,"compositionend"]],function(l,n,u){var t=!0;return"input"===n&&(t=!1!==e["\u0275nov"](l,30)._handleInput(u.target.value)&&t),"blur"===n&&(t=!1!==e["\u0275nov"](l,30).onTouched()&&t),"compositionstart"===n&&(t=!1!==e["\u0275nov"](l,30)._compositionStart()&&t),"compositionend"===n&&(t=!1!==e["\u0275nov"](l,30)._compositionEnd(u.target.value)&&t),t},null,null)),e["\u0275did"](30,16384,null,0,d.DefaultValueAccessor,[e.Renderer2,e.ElementRef,[2,d.COMPOSITION_BUFFER_MODE]],null,null),e["\u0275prd"](1024,null,d.NG_VALUE_ACCESSOR,function(l){return[l]},[d.DefaultValueAccessor]),e["\u0275did"](32,671744,null,0,d.NgModel,[[2,d.ControlContainer],[8,null],[8,null],[2,d.NG_VALUE_ACCESSOR]],{name:[0,"name"],model:[1,"model"]},null),e["\u0275prd"](2048,null,d.NgControl,null,[d.NgModel]),e["\u0275did"](34,16384,null,0,d.NgControlStatus,[d.NgControl],null,null),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275ted"](-1,null,["\n\n          "])),(l()(),e["\u0275eld"](39,0,null,null,17,"div",[["class","row"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275eld"](41,0,null,null,14,"div",[["class","col-sm-12"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](43,0,null,null,11,"div",[["class","form-group"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n                "])),(l()(),e["\u0275eld"](45,0,null,null,1,"label",[["for","name"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["S\u1ed1 \u0111i\u1ec7n tho\u1ea1i ng\u01b0\u1eddi qu\u1ea3n l\xed"])),(l()(),e["\u0275ted"](-1,null,["\n                "])),(l()(),e["\u0275eld"](48,0,null,null,5,"input",[["class","form-control"],["id","name"],["name","phone"],["ngModel",""],["placeholder","Nh\u1eadp s\u1ed1 \u0111i\u1ec7n tho\u1ea1i"],["type","text"]],[[2,"ng-untouched",null],[2,"ng-touched",null],[2,"ng-pristine",null],[2,"ng-dirty",null],[2,"ng-valid",null],[2,"ng-invalid",null],[2,"ng-pending",null]],[[null,"input"],[null,"blur"],[null,"compositionstart"],[null,"compositionend"]],function(l,n,u){var t=!0;return"input"===n&&(t=!1!==e["\u0275nov"](l,49)._handleInput(u.target.value)&&t),"blur"===n&&(t=!1!==e["\u0275nov"](l,49).onTouched()&&t),"compositionstart"===n&&(t=!1!==e["\u0275nov"](l,49)._compositionStart()&&t),"compositionend"===n&&(t=!1!==e["\u0275nov"](l,49)._compositionEnd(u.target.value)&&t),t},null,null)),e["\u0275did"](49,16384,null,0,d.DefaultValueAccessor,[e.Renderer2,e.ElementRef,[2,d.COMPOSITION_BUFFER_MODE]],null,null),e["\u0275prd"](1024,null,d.NG_VALUE_ACCESSOR,function(l){return[l]},[d.DefaultValueAccessor]),e["\u0275did"](51,671744,null,0,d.NgModel,[[2,d.ControlContainer],[8,null],[8,null],[2,d.NG_VALUE_ACCESSOR]],{name:[0,"name"],model:[1,"model"]},null),e["\u0275prd"](2048,null,d.NgControl,null,[d.NgModel]),e["\u0275did"](53,16384,null,0,d.NgControlStatus,[d.NgControl],null,null),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275ted"](-1,null,["\n\n          "])),(l()(),e["\u0275eld"](58,0,null,null,17,"div",[["class","row"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275eld"](60,0,null,null,14,"div",[["class","col-sm-12"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](62,0,null,null,11,"div",[["class","form-group"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n                "])),(l()(),e["\u0275eld"](64,0,null,null,1,"label",[["for","ccnumber"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["M\u1eadt kh\u1ea9u \u0111\u0103ng nh\u1eadp"])),(l()(),e["\u0275ted"](-1,null,["\n                "])),(l()(),e["\u0275eld"](67,0,null,null,5,"input",[["class","form-control"],["id","ccnumber"],["name","password"],["ngModel",""],["placeholder","Nh\u1eadp m\u1eadt kh\u1ea9u \u0111\u0103ng nh\u1eadp"],["type","password"]],[[2,"ng-untouched",null],[2,"ng-touched",null],[2,"ng-pristine",null],[2,"ng-dirty",null],[2,"ng-valid",null],[2,"ng-invalid",null],[2,"ng-pending",null]],[[null,"input"],[null,"blur"],[null,"compositionstart"],[null,"compositionend"]],function(l,n,u){var t=!0;return"input"===n&&(t=!1!==e["\u0275nov"](l,68)._handleInput(u.target.value)&&t),"blur"===n&&(t=!1!==e["\u0275nov"](l,68).onTouched()&&t),"compositionstart"===n&&(t=!1!==e["\u0275nov"](l,68)._compositionStart()&&t),"compositionend"===n&&(t=!1!==e["\u0275nov"](l,68)._compositionEnd(u.target.value)&&t),t},null,null)),e["\u0275did"](68,16384,null,0,d.DefaultValueAccessor,[e.Renderer2,e.ElementRef,[2,d.COMPOSITION_BUFFER_MODE]],null,null),e["\u0275prd"](1024,null,d.NG_VALUE_ACCESSOR,function(l){return[l]},[d.DefaultValueAccessor]),e["\u0275did"](70,671744,null,0,d.NgModel,[[2,d.ControlContainer],[8,null],[8,null],[2,d.NG_VALUE_ACCESSOR]],{name:[0,"name"],model:[1,"model"]},null),e["\u0275prd"](2048,null,d.NgControl,null,[d.NgModel]),e["\u0275did"](72,16384,null,0,d.NgControlStatus,[d.NgControl],null,null),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275ted"](-1,null,["\n\n          "])),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275eld"](79,0,null,null,11,"div",[["class","card-footer"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275eld"](81,0,null,null,3,"button",[["class","btn btn-sm btn-primary"],["type","submit"]],null,[[null,"click"]],function(l,n,u){var t=!0;return"click"===n&&(t=!1!==l.component.onCreateManager(e["\u0275nov"](l,6))&&t),t},null,null)),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275eld"](83,0,null,null,0,"i",[["class","fa fa-dot-circle-o"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,[" L\u01b0u"])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275eld"](86,0,null,null,3,"button",[["class","btn btn-sm btn-danger"],["type","reset"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275eld"](88,0,null,null,0,"i",[["class","fa fa-ban"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,[" Reset"])),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275ted"](-1,null,["\n    "])),(l()(),e["\u0275ted"](-1,null,["\n  "])),(l()(),e["\u0275ted"](-1,null,["\n  "])),(l()(),e["\u0275ted"](-1,null,["\n"])),(l()(),e["\u0275ted"](-1,null,["\n\n"])),(l()(),e["\u0275eld"](96,0,null,null,64,"div",[["class","row"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n  "])),(l()(),e["\u0275eld"](98,0,null,null,61,"div",[["class","col-lg-12"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n    "])),(l()(),e["\u0275eld"](100,0,null,null,58,"div",[["class","card"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n      "])),(l()(),e["\u0275eld"](102,0,null,null,3,"div",[["class","card-header"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275eld"](104,0,null,null,0,"i",[["class","fa fa-align-justify"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,[" Danh s\xe1ch ng\u01b0\u1eddi qu\u1ea3n l\xed \u0111\xe3 \u0111\u0103ng k\xed\n      "])),(l()(),e["\u0275ted"](-1,null,["\n      "])),(l()(),e["\u0275eld"](107,0,null,null,50,"div",[["class","card-body"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275eld"](109,0,null,null,26,"table",[["class","table table-responsive-sm"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275eld"](111,0,null,null,16,"thead",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275eld"](113,0,null,null,13,"tr",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](115,0,null,null,1,"th",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["T\xean ng\u01b0\u1eddi qu\u1ea3n l\xed"])),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](118,0,null,null,1,"th",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["S\u1ed1"])),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](121,0,null,null,1,"th",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["S\u1ed1 \u0111i\u1ec7n tho\u1ea1i"])),(l()(),e["\u0275ted"](-1,null,["\n              "])),(l()(),e["\u0275eld"](124,0,null,null,1,"th",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["Status"])),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275eld"](129,0,null,null,5,"tbody",[],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275and"](16777216,null,null,2,null,g)),e["\u0275did"](132,802816,null,0,o.NgForOf,[e.ViewContainerRef,e.TemplateRef,e.IterableDiffers],{ngForOf:[0,"ngForOf"]},null),e["\u0275pid"](131072,o.AsyncPipe,[e.ChangeDetectorRef]),(l()(),e["\u0275ted"](-1,null,["\n\n          "])),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275eld"](137,0,null,null,19,"ul",[["class","pagination"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275eld"](139,0,null,null,4,"li",[["class","page-item"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275eld"](141,0,null,null,1,"a",[["class","page-link"],["href","#"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["Prev"])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275eld"](145,0,null,null,4,"li",[["class","page-item active"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275eld"](147,0,null,null,1,"a",[["class","page-link"],["href","#"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["1"])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275eld"](151,0,null,null,4,"li",[["class","page-item"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["\n            "])),(l()(),e["\u0275eld"](153,0,null,null,1,"a",[["class","page-link"],["href","#"]],null,null,null,null,null)),(l()(),e["\u0275ted"](-1,null,["Next"])),(l()(),e["\u0275ted"](-1,null,["\n          "])),(l()(),e["\u0275ted"](-1,null,["\n        "])),(l()(),e["\u0275ted"](-1,null,["\n      "])),(l()(),e["\u0275ted"](-1,null,["\n    "])),(l()(),e["\u0275ted"](-1,null,["\n  "])),(l()(),e["\u0275ted"](-1,null,["\n\n"]))],function(l,n){var u=n.component;l(n,32,0,"name",""),l(n,51,0,"phone",""),l(n,70,0,"password",""),l(n,132,0,e["\u0275unv"](n,132,0,e["\u0275nov"](n,133).transform(u.users$)))},function(l,n){l(n,4,0,e["\u0275nov"](n,8).ngClassUntouched,e["\u0275nov"](n,8).ngClassTouched,e["\u0275nov"](n,8).ngClassPristine,e["\u0275nov"](n,8).ngClassDirty,e["\u0275nov"](n,8).ngClassValid,e["\u0275nov"](n,8).ngClassInvalid,e["\u0275nov"](n,8).ngClassPending),l(n,29,0,e["\u0275nov"](n,34).ngClassUntouched,e["\u0275nov"](n,34).ngClassTouched,e["\u0275nov"](n,34).ngClassPristine,e["\u0275nov"](n,34).ngClassDirty,e["\u0275nov"](n,34).ngClassValid,e["\u0275nov"](n,34).ngClassInvalid,e["\u0275nov"](n,34).ngClassPending),l(n,48,0,e["\u0275nov"](n,53).ngClassUntouched,e["\u0275nov"](n,53).ngClassTouched,e["\u0275nov"](n,53).ngClassPristine,e["\u0275nov"](n,53).ngClassDirty,e["\u0275nov"](n,53).ngClassValid,e["\u0275nov"](n,53).ngClassInvalid,e["\u0275nov"](n,53).ngClassPending),l(n,67,0,e["\u0275nov"](n,72).ngClassUntouched,e["\u0275nov"](n,72).ngClassTouched,e["\u0275nov"](n,72).ngClassPristine,e["\u0275nov"](n,72).ngClassDirty,e["\u0275nov"](n,72).ngClassValid,e["\u0275nov"](n,72).ngClassInvalid,e["\u0275nov"](n,72).ngClassPending)})}var v=e["\u0275ccf"]("app-manager",r,function(l){return e["\u0275vid"](0,[(l()(),e["\u0275eld"](0,0,null,null,1,"app-manager",[],null,null,null,p,c)),e["\u0275did"](1,114688,null,0,r,[i.a,a.a,s.k,a.a],null,null)],function(l,n){l(n,1,0)},null)},{},{},[]),m=function(){},f=u("fAE3");u.d(n,"ManagerModuleNgFactory",function(){return C});var C=e["\u0275cmf"](t,[],function(l){return e["\u0275mod"]([e["\u0275mpd"](512,e.ComponentFactoryResolver,e["\u0275CodegenComponentFactoryResolver"],[[8,[v]],[3,e.ComponentFactoryResolver],e.NgModuleRef]),e["\u0275mpd"](4608,o.NgLocalization,o.NgLocaleLocalization,[e.LOCALE_ID,[2,o["\u0275a"]]]),e["\u0275mpd"](4608,d["\u0275i"],d["\u0275i"],[]),e["\u0275mpd"](4608,d.FormBuilder,d.FormBuilder,[]),e["\u0275mpd"](512,o.CommonModule,o.CommonModule,[]),e["\u0275mpd"](512,s.n,s.n,[[2,s.s],[2,s.k]]),e["\u0275mpd"](512,m,m,[]),e["\u0275mpd"](512,d["\u0275ba"],d["\u0275ba"],[]),e["\u0275mpd"](512,d.FormsModule,d.FormsModule,[]),e["\u0275mpd"](512,d.ReactiveFormsModule,d.ReactiveFormsModule,[]),e["\u0275mpd"](512,f.a,f.a,[]),e["\u0275mpd"](512,t,t,[]),e["\u0275mpd"](1024,s.i,function(){return[[{path:"",component:r}]]},[])])})}});