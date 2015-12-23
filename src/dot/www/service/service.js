/**
 * Created by Lin on 2015/8/24.
 */
D.factory('S', function(){
    var factory = {
        expand: function (attr,fn){
            this[attr] = fn();
        }
    };
    return factory;
});

D.run(function ($timeout,S){
    S.expand('mouseenter',function(){
        return function (item,callback,time){
            $timeout.cancel(item.mLeaveTimer);
            item.mEnterTimer = $timeout(function (){
                item.mouseEnterLock = true;
                if(typeof callback === 'function'){
                    callback();
                }
            },time || 100);
        };
    });
});

D.run(function ($timeout,S){
    S.expand('mouseleave',function(){
        return function (item,callback,time){
            $timeout.cancel(item.mEnterTimer);
            item.mLeaveTimer = $timeout(function (){
                item.mouseEnterLock = false;
                if(typeof callback === 'function'){
                    callback();
                }
            },time || 200);
        };
    });
});