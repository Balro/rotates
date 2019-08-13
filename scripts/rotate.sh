#!/usr/bin/env bash
source /etc/profile
source ~/.bash_profile

ROTATE_HOME=`cd -P $(dirname $0)/../;pwd`
export ROTATE_HOME

GO_FILE=$ROTATE_HOME/lib/rotate
CONF_FILE=$ROTATE_HOME/conf/config.toml
LOG_FILE=$ROTATE_HOME/logs/rotate.log
PID_FILE=$ROTATE_HOME/pid

start() {
    pid=`cat $PID_FILE`
    if [[ "$(status)" =~ "OK" ]]; then
        echo "progress is already running in $pid"
        exit 1
    fi
    nohup $GO_FILE $CONF_FILE >> $LOG_FILE 2>&1 &
    echo $! > $PID_FILE
    sleep 3
    status
}

status() {
    pid=`cat $PID_FILE`
    url=`netstat -nlp | grep -w $pid | head -n 1 | awk '{print $4}'`
    exist=`ps -ef | grep -v grep | grep -w $pid | grep -w $GO_FILE | wc -l`
    if [ $exist -gt 0 ];then
        echo "OK, pid $pid"
    else
        echo "progress lost"
    fi
}

stop() {
    pid=`cat $PID_FILE`
    stat=`status`
    if [[ "$stat" =~ "OK" ]]; then
        echo "kill pid $pid"
        kill $pid
    else
        echo "current status $stat"
        exit 0
    fi
    sleep 3
    echo "current status $(status)"
}

usage() {
    echo "$0 <start|stop|status>"
}

case $1 in
    start)
        start
    ;;
    status)
        status
    ;;
    stop)
        stop
    ;;
    *)
        usage
    ;;
esac
