<?php 
	session_start();
	require_once 'php/google-api-php-client/vendor/autoload.php';
?>
<!DOCTYPE html>
<html>
<head>
	<title>Big Query Test</title>
	<meta charset='UTF-8'>
	
	<link href='https://fonts.googleapis.com/css?family=Cabin' rel='stylesheet' type='text/css'>
	<link rel='stylesheet' type='text/css' href='/css/style.css'>
</head>
<body>
	<div id='header'>
		Big Query Result
	</div>
	
	<div class='content'>
	<?php
		$client = new Google_Client();
		$client->useApplicationDefaultCredentials();
		$client->addScope(Google_Service_Bigquery::BIGQUERY);
		$bigquery = new Google_Service_Bigquery($client);
		$projectId = 's3514010-cc2016-go';

		$request = new Google_Service_Bigquery_QueryRequest();
		$str = '';
		
		$request->setQuery("SELECT * FROM [baby.baby_names] LIMIT 10");
		
		$response = $bigquery->jobs->query($projectId, $request);
		$rows = $response->getRows();

		$str = "<table>".
		"<tr>" .
		"<th>Name</th>" .
		"<th>Gender</th>" .
		"<th>Count</th>" .
		"<th>Year</th>" .
		"</tr>";
		
		foreach ($rows as $row)
		{
			$str .= "<tr>";

			foreach ($row['f'] as $field)
			{
				$str .= "<td>" . $field['v'] . "</td>";
			}
			$str .= "</tr>";
		}

		$str .= '</table></div>';

		echo $str;
	?>
	</div>
</body>
</html>
