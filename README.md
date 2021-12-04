
イメージのビルド
```powershell
docker build --rm .\server\brainfuck-runtime\ -t brainfuck-runtime-front
```

起動
```powershell
docker-compose up -d
```

フロント接続
```powershell
docker-compose exec brainfuck-runtime-front bash
```
