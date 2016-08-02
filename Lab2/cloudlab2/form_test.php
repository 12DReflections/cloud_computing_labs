<html>
  <body>
    <?php
    if (array_key_exists('content', $_POST)) {
      echo "You wrote:<pre>\n";
      echo htmlspecialchars($_POST['content']);
      echo "\n</pre>";
    }
    ?>
    <form action="/sign" method="post">
      <div>Write something \n<textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Submit"></div>
    </form>
  </body>
</html>