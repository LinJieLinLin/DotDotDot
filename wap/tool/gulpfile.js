var fs          = require('fs');
var gulp        = require('gulp');
var path        = require('path');
var _           = require('gulp-load-plugins')();
var minimist    = require('minimist');
var C      = require('./config');

var p        = function (dir) {
    return path.resolve(__dirname, dir);
};

var target = [];
var argv   = minimist(process.argv.slice(2));
if (argv.f || argv.file) {
    var files = argv.f || argv.file;
    target = files.split(' ');
} else {
    target = ['*.html'];
}

var files = fs.readdirSync(path.join(__dirname, 'gulp'));

files.forEach(function (file) {
    console.log(file)
    var stat = fs.statSync(path.join(__dirname, 'gulp', file));
    if (stat.isFile() && path.extname(file) === '.js'){
        require(path.join(__dirname, 'gulp', file))(gulp, _, p, C, target);
    }
});

gulp.task('default', ['serve']);