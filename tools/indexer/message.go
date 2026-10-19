package main

const (
	Msg_ProcessingHasFinished = "Processing has finished."
	Msg_BrokenFile            = "Broken file"

	Msgf_InspectedBTFilesCount = "Inspected BitTorrent files: %d."
	Msgf_ProcessingFiles       = "Processing %d BitTorrent files:"
	Msgf_TimeTakenSeconds      = "Time taken: %f seconds."
	Msgf_BytesRead             = "Bytes read: %d."
	Msgf_BytesWritten          = "Bytes written: %d."
	Msgf_FilesIndexed          = "Files indexed: %d."
	Msgf_FilesBroken           = "Broken files: %d."
	Msgf_SelfCheckErrors       = "Self-check errors: %d."
)

const UsageHint = `Usage:
	[ObjectType] [ObjectPath] [Output]

Examples:
	indexer.exe File "123.torrent" "index.csv" 
	indexer.exe Folder "torrents" "index.csv"

Notes:
	Possible object types: File, Folder or Directory.
	Letter case is not important.
	Change directory (CD) to a working directory before usage.`
