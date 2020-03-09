<?php

namespace App\Consumer;

class Article
{
    public static function create()
    {
        sleep(10);
        file_put_contents(APP_ROOT .'/logs/article.log', __CLASS__ .'/'. __FUNCTION__ ."\n", FILE_APPEND);
        echo "create";
    }

    public static function update()
    {
        sleep(10);
        file_put_contents(APP_ROOT .'/logs/article.log', __CLASS__ .'/'. __FUNCTION__ ."\n", FILE_APPEND);
        echo "update";
    }

    public static function delete()
    {
        sleep(10);
        file_put_contents(APP_ROOT .'/logs/article.log', __CLASS__ .'/'. __FUNCTION__ ."\n", FILE_APPEND);
        echo "delete";
    }

    public static function select()
    {
        sleep(10);
        file_put_contents(APP_ROOT .'/logs/article.log', __CLASS__ .'/'. __FUNCTION__ ."\n", FILE_APPEND);
        echo "select";
    }
}