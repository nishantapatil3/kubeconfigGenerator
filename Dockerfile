FROM praqma/network-multitool:latest

# Add nsm-nse code contents
COPY apioverride /apioverride

CMD ["./apioverride"]
