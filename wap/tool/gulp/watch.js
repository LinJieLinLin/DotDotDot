// var browserSync = require('browser-sync');

module.exports = function(gulp, _, dir) {

    // gulp.task('reload', function() {
    //     console.log('刷新数据');
    //     browserSync.reload();
    // });

    gulp.task('watch', function () {
        gulp.watch(dir('../b-start/style/*.scss'), 
            _.sync(gulp).sync(['c-sass']));
        // gulp.watch(dir('../*.js'),
        //     _.sync(gulp).sync(['reload']));
    });
};
