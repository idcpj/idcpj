#!/bin/sh

case $1 in

1)
echo -e "[1]start first server\n"
/usr/bin/etcd --name infra1 \
              --listen-client-urls http://0.0.0.0:2379  \
              --advertise-client-urls http://0.0.0.0:2379 \
              --listen-peer-urls http://0.0.0.0:12380 \
              --initial-advertise-peer-urls http://0.0.0.0:12380 \
              --initial-cluster-token etcd-cluster-1 --initial-cluster \
              'infra1=http://0.0.0.0:12380,infra2=http://0.0.0.0:22380,infra3=http://0.0.0.0:32380' \
              --initial-cluster-state new --enable-pprof
  ;;
2)
echo -e "[2]start second  server\n"
/usr/bin/etcd --name infra2 \
              --listen-client-urls http://0.0.0.0:22379 \
              --advertise-client-urls http://0.0.0.0:22379 \
              --listen-peer-urls http://0.0.0.0:22380 \
              --initial-advertise-peer-urls http://0.0.0.0:22380 \
              --initial-cluster-token etcd-cluster-1 --initial-cluster \
              'infra1=http://0.0.0.0:12380,infra2=http://0.0.0.0:22380,infra3=http://0.0.0.0:32380' \
              --initial-cluster-state new --enable-pprof
  ;;
3)
echo -e "[3]start third server\n"
/usr/bin/etcd --name infra3 \
              --listen-client-urls http://0.0.0.0:32379 \
              --advertise-client-urls http://0.0.0.0:32379 \
              --listen-peer-urls http://0.0.0.0:32380 \
              --initial-advertise-peer-urls http://0.0.0.0:32380 \
              --initial-cluster-token etcd-cluster-1 --initial-cluster \
              'infra1=http://0.0.0.0:12380,infra2=http://0.0.0.0:22380,infra3=http://0.0.0.0:32380' \
              --initial-cluster-state new --enable-pprof
  ;;
*)
echo "error paramater"
;;
esac