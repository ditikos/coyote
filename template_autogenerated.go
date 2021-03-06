package main

const (
mainTemplate = `<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Coyote Tester | Results</title>


    <link rel="stylesheet" href="//ajax.googleapis.com/ajax/libs/angular_material/1.1.0/angular-material.min.css">


    <script src="//use.fontawesome.com/fbbd91a770.js"></script>

    <style>
        body { background-color:#f4f4f4;}
        h1, h2, h3, h4 { font-weight:200 }


        .test-name { width:25%; }
        .test-time { width: 10%; }
        .test-icon { width: 5%; }
        .test-icon,.test-time ,.test-name,.test-command, .test-code  { font-size: 12px;font-weight: bold;color: #333;}
        .test-code { width: 10%; }
        .test-command {width:50%;}
        .main-row { cursor: pointer; }
        .main-row:hover { background-color: #ddd;}
        table {border-spacing: 0;}
        table tbody:nth-child(odd) tr { background-color:#eee; }
        table tbody:nth-child(even) tr { background-color:#fff; }
        table tr td, th {padding:10px;}
        table thead tr { color: rgba(0,0,0,.54); font-size: 12px; font-weight: 700;white-space: nowrap; text-align:left; }
        table tbody { color: rgba(0,0,0,.87);font-size: 13px;vertical-align: middle; }
        table tbody tr { border-top: 1px rgba(0,0,0,.12) solid; }
        .td-hidden {overflow:hidden; padding: 20px;}
        .td-hidden-std {background-color: rgba(244,244,244,0.8);padding:10px;}
        .td-hidden-error {background-color: rgba(212,72,72,0.2);padding:10px;}
        md-card md-card-header md-card-avatar+md-card-header-text  {16px;color:#fff;cursor:pointer;}
        md-content.md-default-theme, md-content {
             color: rgba(0,0,0,0.87);
             background-color: rgb(220,220,220);
        }
        .md-avatar-small { width:10px; height:30px;}
        .icon-header-button {font-size:20px;}
        .icon-status-header-passed {width:20px; margin:10px; font-size:20px; color:green}
        .icon-status-header-failed {width:20px; margin:10px; font-size:20px; color:red}
        .icon-status-passed {width:10px; color:green}
        .icon-status-failed {width:10px; color:red}
        .summary {font-size:14; padding-right:10px;}
        .dark-background {background-color:#2b2b2b; color: #ccc;}
        .logo-section {padding-left:30px;}
        .box {color:#ccc; background-color: #242424;text-align:center; padding: 13px 16px 12px;}
        .execution-details {font-family:"Lucida Console", Monaco, monospace}
        .logo {  margin-bottom: 0px;}
        .summary-section {background-color:#2b2b2b;margin:-20px;}
        .md-button-custom {width:70px; height:70px; margin:10px; background-color:#3f3f3f;}
        .md-button-custom-icon {color:#ccc; size:30px;font-size:20px;}
        .progress-bar {width:100%; -webkit-animation: fullexpand 10s ease-out;}
        .footer-github {text-align:center;margin-bottom:5px;}
        .footer-landoop-img {width:20px;float: left;padding-right:5px;}
         md-card md-card-header md-card-avatar+md-card-header-text .md-title {font-size:18px;}
		.md-headline { font-size: 18px}
		.md-subhead {font-size: 14px;line-height: 24px;}
		.testlist {overflow:auto;height:291px;background:#242424}
		@media (max-width: 1553px) {
		.testlist {height:327px;}

		}
    </style>

    <script src="//cdnjs.cloudflare.com/ajax/libs/d3/3.4.4/d3.min.js"></script>
    <script src="//storage.googleapis.com/artifacts-landoop/d3pie.min.js"></script>
</head>

<body ng-app="CoyoteApp" ng-cloak>

<div ng-controller="MainCtrl">


    <div layout="row" class="dark-background logo-section">
        <div flex="5"></div>
        <div flex layout="column" layout-align="center start">
            <h1 class="logo"><img src="data:image/gif;base64,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" height="40" style="margin-bottom: -5px;"> Coyote | 
                <p class="execution-details md-headline" style="display: inline;">{{datalist.Title}}, <span class="md-subhead">{{datalist.Date}}</span></p>
            </h1>
        </div>
    </div>

    <div layout="row" class="dark-background" style="padding-bottom:30px;">
        <div flex="5"></div>
        <div layout="row" flex layout-margin >
            <div flex="35" layout="column">
                <div class="box">
                    <h4> <b>{{ percentsucc | number:2 }}%</b> Passed</h4>
                    <h6 style="margin-top:-15px;"> Total tests: {{datalist.TotalTests}} </h6>
                </div>

                <md-content class="testlist">
                    <md-list class="md-dense box" flex >


                        <md-list-item ng-repeat="d in datalist.Results" class="md-2-line" ng-click="gotoTest('testNo'+$index)">

                            <i class="" aria-hidden="true" style="margin-right:20px;"
                               ng-class="{ 'fa fa-times icon-status-failed': d.Errors >'0', 'fa fa-check icon-status-passed': d.Errors == '0',  }">
                            </i>
                            <div class="md-list-item-text" layout="column">
                                <h3 style="color:#fff;">{{d.Name}}</h3>
                                <p style="color:#ccc; padding-top: 5px;">Passed {{ d.Passed }} out of {{ d.Total }} | {{d.TotalTime | number:2}}s  </p>
                            </div>
                        <md-divider ></md-divider>
                        </md-list-item>

                    </md-list>
                </md-content>

            </div>
            <div layout="column" flex>
                <div class="box">
                    <h3 style="display:inline-block;">Total duration: <b>{{datalist.TotalTime | number:2}}s</b></h3>
                    <div flex id="testTimes" style="display:inline;"></div>
                </div>
            </div>
        </div>
        <div flex="5"></div>
    </div>

    <div layout="row">
        <div flex="5"></div>
        <div flex>
            <h3>Results</h3>
            <md-card ng-repeat="test in datalist.Results" ng-init="cardIndex = $index" id="testNo{{$index}}"  >
                <md-card-header class="dark-background md-title" ng-click="toggleCard(cardIndex); totalheight()" style="padding-top: 6px; padding-bottom: 6px;">
                    <md-card-avatar layout-align="center start ">
                        <i style="margin-top:10px"
                           ng-class="{ 'fa fa-times icon-status-failed': test.Errors > 0, 'fa fa-check icon-status-passed': test.Errors == 0,  }" aria-hidden="true">
                        </i>
                    </md-card-avatar>
                    <md-card-header-text layout-align="center start">
                        <span class="md-title">{{test.Name}}</span>
                    </md-card-header-text>
                    <md-button class="md-icon-button" aria-label="More">
                        <i  ng-hide="showcard[cardIndex]" class="fa fa-angle-double-up icon-header-button" aria-hidden="true"></i>
                        <i  ng-show="showcard[cardIndex]" class="fa fa-angle-double-down icon-header-button" aria-hidden="true"></i>
                    </md-button>
                </md-card-header>

                <md-content ng-hide="showcard[cardIndex]" >
                    <table style="width:100%;">
                        <thead>
                        <tr>
                            <th></th>
                            <th class="test-icon"><span></span></th>
                            <th class="test-name">Action</th>
                            <th class="test-time"><i class="fa fa-clock-o" aria-hidden="true"></i>  Time (sec)</th>
                            <th hide-sm hide-xs class="test-command"><i class="fa fa-terminal" aria-hidden="true"></i> Command</th>
                            <th class="test-code"><span>Exit code</span></th>
                        </tr>
                        </thead>
                        <tbody  ng-repeat="dtest in test.Results"  ng-init="rowIndex=$index" ng-click="toggleRow(rowIndex,   cardIndex); totalheight();" class="main-row">
                        <tr>
                            <td style="width:10px;">
                                <i class="fa fa-caret-down" aria-hidden="true" ng-hide="showRow[rowIndex+''+cardIndex]"></i>
                                <i class="fa fa-caret-up" aria-hidden="true" ng-show="showRow[rowIndex+''+cardIndex]"></i>
                            </td>
                            <td>
                                <i ng-class="{ 'fa fa-times icon-status-failed': dtest.Status == 'error', 'fa fa-check icon-status-passed': dtest.Status == 'ok',  }" aria-hidden="true"></i>
                            </td>
                            <td><b>{{dtest.Name}}</b></td>
                            <td> {{dtest.Time | number:2}}</td>
                            <td hide-sm hide-xs>
                                <code>
                                    {{dtest.Command}}
                                </code>
                            </td>
                            <td>{{dtest.Exit}}</td>
                        </tr>
                        <tr>
                            <td colspan="6" ng-show="showRow[rowIndex+''+cardIndex]" class="td-hidden">

                                <h4 hide-gt-sm>Command</h4>
                                <div hide-gt-sm>
                                    <code>
                                        {{dtest.Command}}
                                    </code>
                                </div>

                                <h4 ng-show="dtest.Stdout != ''" >Standard Output </h4>
                                <div style="cursor:text" ng-click="$event.stopPropagation();"  ng-show="dtest.Stdout != ''" class="td-hidden-std">
                                    <code>
											<span ng-repeat="stdoutline in dtest.Stdout track by $index" >
												<span ng-bind-html="consoleStdout(stdoutline)" > </span><br />
											</span>
                                    </code>
                                </div>

                                <h4 ng-show="dtest.Stderr != ''" >Standard Error</h4>
                                <div style="cursor:text" ng-click="$event.stopPropagation();"  ng-show="dtest.Stderr != ''" class="td-hidden-error">
                                    <code>
											<span ng-repeat="stderrline in dtest.Stderr track by $index" >
												<span ng-bind-html="consoleStdout(stderrline)" > </span><br />
											</span>
                                    </code>
                                </div>
                            </td>
                        </tr>


                        </tbody>
                    </table>

                    <div layout="row" layout-align="end center">
                        <p class="summary">Passed {{test.Passed}} out of {{test.Total}} | {{test.TotalTime | number:2}} seconds </p>
                    </div>
                </md-content>
            </md-card>
        </div>
        <div flex="5"></div>
    </div>

</div>


<h6 id="#results" class="footer-github">Report Issues & Stars!</h6>
<div flex layout="row" layout-align="center center">
    <a class="github-button" href="https://github.com/Landoop/coyote/issues" data-count-api="/repos/Landoop/coyote#open_issues_count" data-count-aria-label="# issues on GitHub" aria-label="Issue Landoop/coyote on GitHub">Issue</a>
    <a class="github-button" href="https://github.com/Landoop/coyote" data-count-href="/Landoop/coyote/stargazers" data-count-api="/repos/Landoop/coyote#stargazers_count" data-count-aria-label="# stargazers on GitHub" aria-label="Star Landoop/coyote on GitHub">Star</a>
</div>

<div flex layout="row" layout-align="center center">
    <img ng-src="https://www.landoop.com/images/landoop-blue.svg" class="footer-landoop-img">
    <p style="font-size:10px;">powered by <a href="https://www.landoop.com" style="text-decoration:none;color:blue;" target="_blank">Landoop</a></p>
</div>


<script src="//ajax.googleapis.com/ajax/libs/angularjs/1.5.5/angular.min.js"></script>
<script src="//ajax.googleapis.com/ajax/libs/angularjs/1.5.5/angular-animate.min.js"></script>
<script src="//ajax.googleapis.com/ajax/libs/angularjs/1.5.5/angular-aria.min.js"></script>
<script src="//ajax.googleapis.com/ajax/libs/angularjs/1.5.5/angular-messages.min.js"></script>
<script src="//ajax.googleapis.com/ajax/libs/angularjs/1.2.0rc1/angular-route.min.js"></script>



<script src="//ajax.googleapis.com/ajax/libs/angular_material/1.1.0/angular-material.min.js"></script>
<script src="//storage.googleapis.com/wch/ansi_up-1.3.0.js" type="text/javascript"></script>


<script async defer src="//buttons.github.io/buttons.js"></script>


<script type="text/javascript">

     angular.module('CoyoteApp', ['ngMaterial', 'ngAnimate', 'ngAria'])
            .controller('MainCtrl', function ($scope, $log, $location, $anchorScroll, $sce) {

            var data = <{=( .Data )=}> ;

	        function getRandomColor() {
		    var letters = '0123456789ABCDEF';
		    var color = '#';
                for (var i = 0; i < 6; i++ ) {
                color += letters[Math.floor(Math.random() * 16)];
                }
		    return color;
	        }

	        $scope.gotoTest = function(testid) {
	            var oldlocation = $location.hash();
	            $location.hash(testid);
                $anchorScroll();
                //reset to old to keep any additional routing logic from kicking in
                $location.hash(oldlocation);
	        }

	        $scope.consoleStdout = function(stdout) {
		    var ansiStdout= ansi_up.ansi_to_html(stdout);
		    var trustedAnsiStdout = $sce.trustAsHtml(ansiStdout)
		        return 	trustedAnsiStdout;
	        }

	        $scope.datalist = data;

	        $scope.percentsucc = data.Successful / data.TotalTests * 100;
	        document.title = "Coyote Tester | " + $scope.datalist.Title + " | Results";

	        $scope.showcard=[]
	        $scope.toggleCard = function(id) {
		    $scope.showcard[id] = !$scope.showcard[id];
	        }

	        $scope.showRow=[]
	        $scope.toggleRow = function(  rowIndex,  cardIndex) {
		    $scope.showRow[rowIndex+""+cardIndex] = !$scope.showRow[rowIndex+""+cardIndex];
	        }

	        $scope.toggleDetail = function($index) {
	            $scope.activePosition = $scope.activePosition == $index ? -1 : $index;
	        };
            var colorpallete = [ "#2383c1","#64a61f","#7b6788","#a05c56","#961919","#d8d239","#e98125","#d0743c","#635122","#6ada6a","#0b6197","#7c9058","#207f32","#44b9af","#bca349",]
	        var content = [];
	        var i=0;

	        angular.forEach($scope.datalist.Results, function(results2) {
		    angular.forEach(results2.Results, function(results3) {
		        this.push({"label": results2.Name + results3.Name  ,"value": results3.Time,"color": colorpallete[i]});
		        i++;
		    }, content);
	        });

			function findTotalheight() {
				parent.postMessage(document.body.scrollHeight,"*");
			}

			$scope.totalheight = function() {
				setTimeout(findTotalheight, 200)
			}

	        var pie = new d3pie("testTimes", {
		    "header": {
			"title": {
			    "text": "",
			    "color": "#fffefe",
			    "fontSize": 34,
			    "font": "sans"
		        },
			"subtitle": {
			    "text": "",
			    "color": "#999999",
			    "fontSize": 14,
			    "font": "sans"
			},
			"location": "pie-center",
			"titleSubtitlePadding": 10
		    },
		    "footer": {
			"text": "",
			"color": "#999999",
			"fontSize": 10,
			"font": "open sans",
			"location": "bottom-left"
		    },
		    "size": {
			"canvasHeight": 350,
			"canvasWidth": 650,
			"pieInnerRadius": "72%",
			"pieOuterRadius": "85%"
		    },
		    "data": {
			"sortOrder": "label-desc",
			"smallSegmentGrouping": {
			    "enabled": true,
			    "value": 3
			},
		        "content": content
		    },
		    "labels": {
			"outer": {
			    "pieDistance": 25
			},
			"inner": {
			    "hideWhenLessThanPercentage": 3
			},
			"mainLabel": {
			    "color": "#ffffff",
			    "fontSize": 10,
			    "pieDistance": 15,
			    "padding": 4
			},
			"value": {
			    "color": "#cccc43",
			    "fontSize": 10
			},
			"lines": {
			    "enabled": true,
			    "color": "#777777"
			},
			"truncation": {
			    "enabled": true
			}
		    },
		    "effects": {
			"pullOutSegmentOnClick": {
			    "effect": "linear",
			    "speed": 400,
			    "size": 8
			}
		    },
		    "misc": {
			"colors": {
			    "background": "#242424",
			    "segmentStroke": "#f6f6f6"
			}
		    }
	        });
            });
    </script>

</body>
<{=( .Version )=}>
</html>
`
)
