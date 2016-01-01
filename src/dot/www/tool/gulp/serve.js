// var notify = require("gulp-notify");
module.exports = function(gulp, _, dir) {
    gulp.task('default', function() {
        // .pipe(notify("Hello Gulp!"));
        console.log('hello world');
    });

    gulp.task('serve', _.sync(gulp).sync([
        // 'copyimage', 
        // 'directive:css', 
        // 'directive:tmpl', 
        'compile:css', 
        // 'compile:js', 
        // 'compile:html', 
        // 'proxyServer', 
        'watch']), function () {
        console.log('hello world');
    });
};
