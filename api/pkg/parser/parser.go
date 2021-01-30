package parser

import (
	"archive/zip"
	"bytes"
	"errors"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/andrianbdn/iospng"
	"github.com/appflight/apkparser"
	"howett.net/plist"
)

var (
	reInfoPlist = regexp.MustCompile(`Payload/[^/]+/Info\.plist`)
	ErrNoIcon   = errors.New("icon not found")
)

const (
	iosExt          = ".ipa"
	androidExt      = ".apk"
	PlatformIos     = 1
	PlatformAndroid = 2
)

type AppInfo struct {
	Name     string
	BundleId string
	Version  string
	Build    string
	Platform uint8
	Icon     image.Image
	Size     uint64
}

type androidManifest struct {
	Package     string `xml:"package,attr"`
	VersionName string `xml:"versionName,attr"`
	VersionCode string `xml:"versionCode,attr"`
}

type iosPlist struct {
	CFBundleName         string `plist:"CFBundleName"`
	CFBundleDisplayName  string `plist:"CFBundleDisplayName"`
	CFBundleVersion      string `plist:"CFBundleVersion"`
	CFBundleShortVersion string `plist:"CFBundleShortVersionString"`
	CFBundleIdentifier   string `plist:"CFBundleIdentifier"`
}

func NewAppParser(name string) (*AppInfo, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	reader, err := zip.NewReader(file, stat.Size())
	if err != nil {
		return nil, err
	}

	var plistFile, iosIconFile *zip.File
	for _, f := range reader.File {
		switch {
		case reInfoPlist.MatchString(f.Name):
			plistFile = f
		case strings.Contains(f.Name, "AppIcon"):
			iosIconFile = f
		}
	}

	ext := filepath.Ext(stat.Name())

	if ext == androidExt {
		apkInfo, err := apkparser.ParseApk(name)
		if err != nil {
			logrus.Error("parseApkIconAndLabel err:", err)
		}
		var info AppInfo
		info.Name = apkInfo.Label
		info.BundleId = apkInfo.Package
		info.Version = apkInfo.VersionName
		info.Build = strconv.FormatInt(int64(apkInfo.VersionCode), 10)
		info.Platform = PlatformAndroid
		info.Icon = apkInfo.Icon
		info.Size = uint64(stat.Size())
		return &info, err
	}

	if ext == iosExt {
		info, err := parseIpaFile(plistFile)
		icon, err := parseIpaIcon(iosIconFile)
		info.Icon = icon
		info.Size = uint64(stat.Size())
		info.Platform = PlatformIos
		return info, err
	}

	return nil, errors.New("unknown platform")
}

func parseIpaFile(plistFile *zip.File) (*AppInfo, error) {
	if plistFile == nil {
		return nil, errors.New("info.plist not found")
	}

	rc, err := plistFile.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	buf, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	p := new(iosPlist)
	decoder := plist.NewDecoder(bytes.NewReader(buf))
	if err := decoder.Decode(p); err != nil {
		return nil, err
	}

	info := new(AppInfo)
	if p.CFBundleDisplayName == "" {
		info.Name = p.CFBundleName
	} else {
		info.Name = p.CFBundleDisplayName
	}
	info.BundleId = p.CFBundleIdentifier
	info.Version = p.CFBundleShortVersion
	info.Build = p.CFBundleVersion

	return info, nil
}

func parseIpaIcon(iconFile *zip.File) (image.Image, error) {
	if iconFile == nil {
		return nil, ErrNoIcon
	}

	rc, err := iconFile.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	var w bytes.Buffer
	iospng.PngRevertOptimization(rc, &w)

	return png.Decode(bytes.NewReader(w.Bytes()))
}
