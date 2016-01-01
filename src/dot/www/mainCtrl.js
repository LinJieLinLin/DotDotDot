D.controller('mainCtrl', ['$scope', '$timeout', 'S', function($scope, $timeout, S) {
    $scope.hi = 'hello World';
    $scope.getData = function(params, filter) {
        var option = {
            method: 'GET',
            url: '/getUser',
            params: params
        };
        return S.R(angular.extend(option, filter || {}));
    };
    // $scope.getData({a:1});
    $scope.tabIndex = 0;
    $scope.tabs = [{
        name: '姓名'
    }, {
        name: '菜单'
    }];
}]);
