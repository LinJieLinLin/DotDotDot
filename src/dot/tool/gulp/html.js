module.exports = function (gulp, _, dir, config, target) {
    // inject bower's dependencies to html files
    function injectBowerFile () {
        return _.inject(gulp.src(dir('../www/application/module/common/vendor*'), {read: false}), {
            name: 'bower',
            transform: function (filepath, file, i, length, target) {
                // /../www/
                filepath    = filepath.substring(8); 
                return _.inject.transform.call(_.inject.transform, filepath, file, i, length, target);
            }
        });
    }

    function injectCommon () {
        var source = [];
        config.common.forEach(function (item) {
            source.push(dir('dist/' + item));
        });
        return _.inject(gulp.src(source, {read: false}), {
            name: 'common',
            transform: function (filepath, file, index, length, targetFile) {
                filepath = filepath.substring(6);
                return _.inject.transform.call(_.inject.transform, filepath, file, index, length, targetFile);
            }
        });
    }

    function injectLoading () {
        return _.inject(gulp.src(dir('gulpfile.js'), {read: false}), {
            name: 'loading',
            transform: function () {
                // var script = ['<script>window.loadingLayer.in();window.addEventListener("load",function(){this.loadingLayer.out()},false);</script>'].join('');
                var script = [''].join('');
                return script;
            }
        });
    }

    function injectThirdPartyImport(){
        return _.inject(gulp.src(dir('gulpfile.js'), {read: false}), {
            name: 'thirdPartyImport',
            transform: function (){
                var script = [
                    // '<script type="text/javascript" src="http://qzonestyle.gtimg.cn/qzone/openapi/qc_loader.js" data-appid="101254692" data-redirecturi="http://sso.t.jxzy.com" charset="utf-8" data-callback="true"></script>',
                    // '<script src="http://tjs.sjs.sinajs.cn/open/api/js/wb.js?appkey=886838876&debug=true" type="text/javascript" charset="utf-8"></script>',
                    '<script src="http://s11.cnzz.com/z_stat.php?id=1256639613&web_id=1256639613" language="JavaScript"></script>'
                ].join('');
                return script;
            }
        });
    }

    gulp.task('compile:html', function () {
        var t = [];

        target.forEach(function (item) {
            t.push(dir('src/' + item));
        });

        return gulp.src(t)
        .pipe(_.changed(dir('dist')))
        //add ngCloak
        .pipe(_.angularCloak())
        // replace the assets' placeholder to the actual code
        .pipe(injectBowerFile())
        .pipe(injectCommon())
        .pipe(injectLoading())
        .pipe(injectThirdPartyImport())
        
        .pipe(gulp.dest(dir('dist')));
    }); 
};