until myslq -h db -u root -p db_root_password &> /dev/null
do
    echo 'waiting for myslqd to be connectable'
    sleep 2
done

echo "app is starting...!"
exec go run main.go