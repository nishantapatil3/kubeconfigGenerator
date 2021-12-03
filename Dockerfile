FROM praqma/network-multitool:latest

COPY kubeconfigGenerator /kubeconfigGenerator

CMD ["./kubeconfigGenerator"]
