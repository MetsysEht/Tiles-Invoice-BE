# Define the base directory for proto files and output directory
$protoPath = "././proto"
$outPath = "././rpc"

# Delete the rpc folder if it exists, then recreate it
if (Test-Path $outPath) {
    Remove-Item -Recurse -Force $outPath
}
New-Item -ItemType Directory -Path $outPath

# Get all .proto files recursively from the proto directory
$protoFiles = Get-ChildItem -Recurse -Filter *.proto -Path $protoPath

# Loop through each .proto file and run protoc
foreach ($file in $protoFiles) {
    # Get the relative path of the proto file by removing the base path
    $relativeFilePath = $file.FullName.Substring((Get-Location).Path.Length + 1)

    # Run the protoc command with the correct relative file path
    protoc --proto_path=$protoPath --go_out=$outPath --go-grpc_out=$outPath --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $relativeFilePath
}
