apt-get update && \
    apt-get install -y pkg-config gcc libseccomp-dev && \
    apt-get clean && rm -rf /var/lib/apt/lists/*
