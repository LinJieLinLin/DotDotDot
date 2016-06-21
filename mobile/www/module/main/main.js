var dot = ons.bootstrap();
dot.controller('main', ['$scope', function($scope) {
    this.load = function(page) {
        dot.content.load(page)
            .then(function() {
                dot.left.close();
            });
    };
    $scope.sb = '1234';
}]);
