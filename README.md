# go-walvis
Go playground for data visualization

## 1 Roadmaps
- (Util) List all csv files.
- DataFrame (from csv) & Analysis
    - Total expense per account per month
    - Average expense per category per month (PIE Chart)
    - Total expense per month (PIE Chart)
- Visualiser
- API server (receiver csv files)

## 2 Development

### 2.1 Build and run a Docker container
```
cd ./go-walvis
docker build -f docker/Dockerfile -t go-walvis .
docker run -it --rm -v $PWD:/workspace go-walvis /bin/bash
```

### 2.2 Build main executable

```
go build cmd/walvis/main.go
```

### 2.3 Test a module

```
cd ./walvis
go test
```
