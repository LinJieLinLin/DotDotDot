var browserSync = require('browser-sync');
var path = require('path');
module.exports = function(gulp, _, dir) {
    gulp.task('default', function() {
        console.log('default:hello world');
    });
    gulp.task('proxyServer', function() {
        gulp.src(dir('proxy-server.js'))
            .pipe(_.shell('node <%= file.path %>'));
    });

    gulp.task('serve', _.sync(gulp).sync([
        // 'copyimage', 
        // 'directive:css', 
        // 'directive:tmpl', 
        'compile:css',
        // 'compile:js', 
        // 'compile:html', 
        'proxyServer',
        'watch'
    ]), function() {
        console.log('serve:hello world');
        browserSync.init({
            proxy: {
                target: 'http://localhost:8888',
                reqHeaders: function(config) {
                    return {
                        agent: false
                    };
                }
            },
            port: 9000,
            ghostMode: false
        });
    });
};
